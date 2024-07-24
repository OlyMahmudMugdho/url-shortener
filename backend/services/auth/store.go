package auth

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/OlyMahmudMugdho/url-shortener/types"
)

type AuthStore struct {
	db *sql.DB
}

func NewAuthStore(db *sql.DB) *AuthStore {
	return &AuthStore{
		db: db,
	}
}

func (a *AuthHandler) SaveUser(user types.User) error {

	var query string = fmt.Sprintf(`INSERT INTO users (USERNAME, PASSWORD, FIRST_NAME, LAST_NAME, EMAIL) VALUES ('%s','%s','%s','%s','%s')`, user.Username, user.Password, user.FirstName, user.LastName, user.LastName)

	_, error := a.store.db.Exec(query)

	if error != nil {
		log.Fatal(error)
		return error
	}
	return nil
}
