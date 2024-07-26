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

func (h *AuthStore) SaveUser(user models.User) error {

	var query string = fmt.Sprintf(`INSERT INTO users (USERNAME, PASSWORD, FIRST_NAME, LAST_NAME, EMAIL) VALUES ('%s','%s','%s','%s','%s')`, user.Username, user.Password, user.FirstName, user.LastName, user.Email)

	_, error := h.db.Exec(query)

	if error != nil {
		return error
	}
	return nil
}

func (h *AuthStore) FindUserByUserName(username string) (models.User, error) {
	user := new(models.User)
	var query string = `SELECT * FROM "users" WHERE USERNAME=$1`
	row := h.db.QueryRow(query, username)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Email)

	if err != nil {
		return *user, err
	}

	return *user, nil
}
