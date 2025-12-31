package utils

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/OlyMahmudMugdho/url-shortener/models"
	"github.com/golang-jwt/jwt/v5"

	"crypto/sha256"
	"encoding/binary"

	"github.com/OlyMahmudMugdho/url-shortener/types"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func ConnectToDatabase() (*sql.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = "5432"
	}

	var config = types.PostgresConfig{
		Username: os.Getenv("POSTGRES_USERNAME"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Db:       os.Getenv("POSTGRES_DBNAME"),
		Sslmode:  os.Getenv("POSTGRES_SSLMODE"),
		Host:     host,
		Port:     port,
	}
	var connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.Db, config.Sslmode)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Println("Error opening database:", err)
		return nil, err
	} else {
		return db, nil
	}
}

func CreateTables(db *sql.DB) error {
	userSchemaFile, err := os.ReadFile("db/user/user_table_up.sql")
	if err != nil {
		fmt.Println(err)
		return err
	}

	urlSchemaFile, err := os.ReadFile("db/shortener/url_table_up.sql")
	if err != nil {
		fmt.Println(err)
		return err
	}

	userSchema := string(userSchemaFile)
	urlSchema := string(urlSchemaFile)

	_, err = db.Exec(userSchema)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = db.Exec(urlSchema)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

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

func GetUserIdFromContext(ctx context.Context) string {
	var username types.ContextKey = "userId"
	return ctx.Value(username).(string)
}

func ExtractParamFromUrl(url string, prefix string) (string, bool) {
	linkId, ok := strings.CutPrefix(url, prefix)
	return linkId, ok
}

func DbErrorMessage(err error, entityName string) string {
	pgErr, _ := err.(*pq.Error)
	errorName := pgErr.Code.Name()

	if errorName == "unique_violation" {
		return entityName + " already exists"
	} else {
		return "something went wrong"
	}
}

func GenerateShortUrl(url string) string {
	hash := sha256.Sum256([]byte(url + time.Now().String()))
	// Take first 8 bytes and convert to uint64
	num := binary.BigEndian.Uint64(hash[:8])
	return Base62Encode(num)
}

func Base62Encode(n uint64) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if n == 0 {
		return string(alphabet[0])
	}
	var res []byte
	for n > 0 {
		res = append(res, alphabet[n%62])
		n /= 62
	}
	// Reverse the slice
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	// Return only first 7 characters for shorter URLs
	if len(res) > 7 {
		return string(res[:7])
	}
	return string(res)
}
