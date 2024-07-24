package utils

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/OlyMahmudMugdho/url-shortener/types"
)

func ConnectToDatabase() bool {
	var config types.PostgresConfig = types.PostgresConfig{
		Username: os.Getenv("POSTGRES_USERNAME"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Db:       os.Getenv("POSTGRES_DBNAME"),
		Sslmode:  os.Getenv("POSTGRES_SSLMODE"),
	}
	var connStr string = "user=" + config.Username + " password=" + config.Password + " dbname=" + config.Db + " sslmode=" + config.Sslmode
	_, error := sql.Open("postgres", connStr)

	if error != nil {
		fmt.Println(error)
		return false
	} else {
		return true
	}
}
