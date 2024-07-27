package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/OlyMahmudMugdho/gotenv/gotenv"
	"github.com/OlyMahmudMugdho/url-shortener/middlewares"
	"github.com/OlyMahmudMugdho/url-shortener/services/auth"
	"github.com/OlyMahmudMugdho/url-shortener/types"
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
	h.router.Handle("GET /dev", middlewares.VerifyAuthentication(Hello))
	authHandler.RegisterRoutes(h.router)

	log.Printf("server is listening on port %v", h.port)
	error := http.ListenAndServe(h.port, h.router)

	if error != nil {
		log.Fatal(error)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	var k types.ContextKey = "username"
	ctx := r.Context().Value(k)
	fmt.Println(ctx)
	w.Write([]byte(ctx.(string)))
}
