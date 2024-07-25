package api

import (
	"database/sql"
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

	db, _ := utils.ConnectToDatabase()

	router := http.NewServeMux()

	return &ApiServer{
		port:   port,
		db:     db,
		router: router,
	}
}
