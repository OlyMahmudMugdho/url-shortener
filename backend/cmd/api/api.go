package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/OlyMahmudMugdho/gotenv/gotenv"
	"github.com/OlyMahmudMugdho/url-shortener/services/auth"
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

func (h *ApiServer) Run() {

	err := h.db.Ping()

	if err != nil {
		log.Fatal(`error connecting to database`)
	} else {
		log.Println("connected to database")
	}

	authStore := auth.NewAuthStore(h.db)
	authHandler := auth.NewAuthHandler(authStore)
	authHandler.RegisterRoutes(h.router)

	log.Printf("server is listening on port %v", h.port)
	error := http.ListenAndServe(h.port, h.router)

	if error != nil {
		log.Fatal(error)
	}
}
