package auth

import (
	"log"
	"net/http"

	"github.com/OlyMahmudMugdho/url-shortener/types"
)

type AuthHandler struct {
	store *AuthStore
}

func NewAuthHandler(store *AuthStore) *AuthHandler {
	return &AuthHandler{
		store: store,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request, userBody types.UserRequestBody) any {
	user := new(types.User)

	user.Username = userBody.Username
	user.Password = userBody.Password
	user.Email = userBody.Email
	user.FirstName = userBody.FirstName
	user.LastName = userBody.LastName

	error := h.SaveUser(*user)

	if error != nil {
		log.Fatal(error)
		w.WriteHeader(403)
		return nil
	} else {
		return userBody
	}
}
