package auth

import (
	"database/sql"

	"github.com/OlyMahmudMugdho/url-shortener/models"
)

type Store struct {
	db *sql.DB
}

func NewAuthStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (h *Store) SaveUser(user models.User) error {

	var query = `INSERT INTO users (USERNAME, PASSWORD, FIRST_NAME, LAST_NAME, EMAIL) VALUES ($1, $2, $3, $4, $5)`

	_, err := h.db.Exec(query, user.Username, user.Password, user.FirstName, user.LastName, user.Email)

	if err != nil {
		return err
	}
	return nil
}

func (h *Store) FindUserByUserName(username string) (models.User, error) {
	user := new(models.User)
	var query = `SELECT * FROM "users" WHERE USERNAME=$1`
	row := h.db.QueryRow(query, username)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Email)

	if err != nil {
		return *user, err
	}

	return *user, nil
}
