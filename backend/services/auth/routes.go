package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/OlyMahmudMugdho/url-shortener/types"
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

	userBody := new(types.UserRequestBody)
	json.NewDecoder(r.Body).Decode(&userBody)

	user := utils.GetUserFromUserRequest(userBody)

	hashedPassword, err := utils.HashPassword(userBody.Password)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500) // internal server error
	}

	user.Password = string(hashedPassword)

	error := h.SaveUser(user)

	if error != nil {
		log.Println(error)
		w.WriteHeader(403) // unauthenticated
	} else {
		json.NewEncoder(w).Encode(utils.GenerateUserResponseFromUser(user))
	}
}
