package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/OlyMahmudMugdho/gotenv/gotenv"
	"github.com/OlyMahmudMugdho/url-shortener/utils"
)

type ApiServer struct {
	port   string
	router *http.ServeMux
	db     *sql.DB
}

func NewApiServer(port string) *ApiServer {
	gotenv.Load()

	db, error := utils.ConnectToDatabase()

	if error != nil {
		log.Fatal(`error connecting to database`)
		return nil
	}

	router := http.NewServeMux()

	return &ApiServer{
		port:   port,
		db:     db,
		router: router,
	}
}
