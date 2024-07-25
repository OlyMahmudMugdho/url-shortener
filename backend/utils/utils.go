package utils

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/OlyMahmudMugdho/url-shortener/types"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func ConnectToDatabase() (*sql.DB, error) {
	var config types.PostgresConfig = types.PostgresConfig{
		Username: os.Getenv("POSTGRES_USERNAME"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Db:       os.Getenv("POSTGRES_DBNAME"),
		Sslmode:  os.Getenv("POSTGRES_SSLMODE"),
	}
	var connStr string = "user=" + config.Username + " password=" + config.Password + " dbname=" + config.Db + " sslmode=" + config.Sslmode
	db, error := sql.Open("postgres", connStr)

	if error != nil {
		fmt.Println(error)
		return nil, error
	} else {
		return db, nil
	}
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 10)
}

func IsPassWordValid(password string, hash string) bool {
	error := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return error == nil
}
