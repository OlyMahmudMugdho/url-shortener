package utils

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/OlyMahmudMugdho/url-shortener/models"

	"github.com/OlyMahmudMugdho/url-shortener/types"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func ConnectToDatabase() (*sql.DB, error) {
	var config = types.PostgresConfig{
		Username: os.Getenv("POSTGRES_USERNAME"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Db:       os.Getenv("POSTGRES_DBNAME"),
		Sslmode:  os.Getenv("POSTGRES_SSLMODE"),
	}
	var connStr = "user=" + config.Username + " password=" + config.Password + " dbname=" + config.Db + " sslmode=" + config.Sslmode
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return db, nil
	}
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 10)
}

func IsPassWordValid(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateUserResponseFromUser(user models.User) models.UserResponseBody {
	userResponse := new(models.UserResponseBody)

	userResponse.Username = user.Username
	userResponse.Email = user.Email
	userResponse.FirstName = user.FirstName
	userResponse.LastName = user.LastName

	return *userResponse
}

func GetUserFromUserRequest(userRequest *models.UserRequestBody) models.User {
	user := new(models.User)

	user.Username = userRequest.Username
	user.Password = userRequest.Password
	user.Email = userRequest.Email
	user.FirstName = userRequest.FirstName
	user.LastName = userRequest.LastName

	return *user
}
