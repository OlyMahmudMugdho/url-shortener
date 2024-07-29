package utils

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/OlyMahmudMugdho/url-shortener/models"
	"github.com/golang-jwt/jwt/v5"

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

func GenerateJWT(username string, userId string) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"username": username,
		"userId":   userId,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(secret))

	return tokenStr, err
}

func ExtractToken(cookieName string, cookies []*http.Cookie) string {
	var token string
	for _, v := range cookies {
		if v.Name == "token" {
			token = v.Value
		}
	}
	return token
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid signing method : %v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

func GetValueFromContext(ctx context.Context, key types.ContextKey) string {
	val := ctx.Value(key).(string)
	return val
}

func GetUsernameFromContext(ctx context.Context) string {
	var username types.ContextKey = "username"
	return ctx.Value(username).(string)
}
