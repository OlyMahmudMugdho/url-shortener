package auth

import (
	"database/sql"
	"fmt"
	"github.com/OlyMahmudMugdho/url-shortener/models"
)

type AuthStore struct {
	db *sql.DB
}

func NewAuthStore(db *sql.DB) *AuthStore {
	return &AuthStore{
		db: db,
	}
}

func (h *AuthHandler) SaveUser(user models.User) error {

	var query string = fmt.Sprintf(`INSERT INTO users (USERNAME, PASSWORD, FIRST_NAME, LAST_NAME, EMAIL) VALUES ('%s','%s','%s','%s','%s')`, user.Username, user.Password, user.FirstName, user.LastName, user.Email)

	_, error := h.store.db.Exec(query)

	if error != nil {
		return error
	}
	return nil
}
