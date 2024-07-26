package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/OlyMahmudMugdho/url-shortener/models"

	"github.com/OlyMahmudMugdho/url-shortener/utils"
)

type AuthHandler struct {
	store *AuthStore
}

func NewAuthHandler(store *AuthStore) *AuthHandler {
	return &AuthHandler{
		store: store,
	}
}

func (h *AuthHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /register", h.Register)
	router.HandleFunc("POST /login", h.Login)
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {

	userBody := new(models.UserRequestBody)
	err := json.NewDecoder(r.Body).Decode(&userBody)
	if err != nil {
		return
	}

	user := utils.GetUserFromUserRequest(userBody)

	hashedPassword, err := utils.HashPassword(userBody.Password)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500) // internal server err2
	}

	user.Password = string(hashedPassword)

	err2 := h.store.SaveUser(user)

	if err2 != nil {
		log.Println(err2)
		w.WriteHeader(403) // unauthenticated
	} else {
		err := json.NewEncoder(w).Encode(utils.GenerateUserResponseFromUser(user))
		if err != nil {
			return
		}
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest := new(models.LoginRequest)

	json.NewDecoder(r.Body).Decode(&loginRequest)

	user, err := h.store.FindUserByUserName(loginRequest.Username)

	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		return
	}

	var validPassword bool = utils.IsPassWordValid(loginRequest.Password, user.Password)

	if !validPassword {
		w.WriteHeader(403)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid credentials",
		})
		return
	}

	token, tokenErr := utils.GenerateJWT(user.Username)

	if tokenErr != nil {
		log.Println(tokenErr)
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "error generating token",
		})
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, cookie)

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(models.LoginResponse{
		Token: token,
	})
}
