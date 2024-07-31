package auth

import (
	"encoding/json"
	"github.com/OlyMahmudMugdho/url-shortener/types"
	"log"
	"net/http"

	"github.com/OlyMahmudMugdho/url-shortener/models"

	"github.com/OlyMahmudMugdho/url-shortener/utils"
)

type Handler struct {
	store *Store
}

func NewAuthHandler(store *Store) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /register", h.Register)
	router.HandleFunc("POST /login", h.Login)
	router.HandleFunc("GET /logout", h.LogOut)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

	userBody := new(models.UserRequestBody)
	err := json.NewDecoder(r.Body).Decode(&userBody)
	if err != nil {
		return
	}

	user := utils.GetUserFromUserRequest(userBody)

	hashedPassword, hashError := utils.HashPassword(userBody.Password)

	if hashError != nil {
		log.Println(hashError)
		w.WriteHeader(500) // internal server error
		json.NewEncoder(w).Encode(&types.Error{
			Error:   true,
			Message: "internal server error",
			Code:    http.StatusInternalServerError,
		})
	}

	user.Password = string(hashedPassword)

	saveError := h.store.SaveUser(user)

	if saveError != nil {
		log.Println(saveError)
		message := utils.DbErrorMessage(saveError, "user")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&types.Error{
			Error:   true,
			Message: message,
			Code:    http.StatusBadRequest,
		}) // unauthenticated
		return
	}
	encodeError := json.NewEncoder(w).Encode(utils.GenerateUserResponseFromUser(user))
	if encodeError != nil {
		return
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest := new(models.LoginRequest)

	decodingError := json.NewDecoder(r.Body).Decode(&loginRequest)

	if decodingError != nil {
		w.WriteHeader(500)
	}

	user, err := h.store.FindUserByUserName(loginRequest.Username)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&types.Error{
			Error:   true,
			Message: "user not found with this username",
			Code:    http.StatusNotFound,
		})
		return
	}

	var validPassword = utils.IsPassWordValid(loginRequest.Password, user.Password)

	if !validPassword {
		w.WriteHeader(403)
		err := json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid credentials",
		})
		if err != nil {
			return
		}
		return
	}

	token, tokenErr := utils.GenerateJWT(user.Username, user.Id)

	if tokenErr != nil {
		log.Println(tokenErr)
		w.WriteHeader(500)
		err := json.NewEncoder(w).Encode(map[string]string{
			"error": "error generating token",
		})
		if err != nil {
			return
		}
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
	err = json.NewEncoder(w).Encode(models.LoginResponse{
		Token: token,
	})
	if err != nil {
		return
	}
}

func (h *Handler) LogOut(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "token",
		Value: "",
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(200)
}
