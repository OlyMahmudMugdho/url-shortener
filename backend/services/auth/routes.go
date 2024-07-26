package auth

import (
	"encoding/json"
	"github.com/OlyMahmudMugdho/url-shortener/models"
	"log"
	"net/http"

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

	err2 := h.SaveUser(user)

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
