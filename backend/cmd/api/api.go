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

type Server struct {
	port   string
	router *http.ServeMux
	db     *sql.DB
}

func NewApiServer(port string) *Server {
	err := gotenv.Load()
	if err != nil {
		return nil
	}

	db, _ := utils.ConnectToDatabase()

	router := http.NewServeMux()

	return &Server{
		port:   port,
		db:     db,
		router: router,
	}
}

func (h *Server) Run() {

	err := h.db.Ping()

	if err != nil {
		log.Fatal(`err2 connecting to database`)
	} else {
		log.Println("connected to database")
	}

	authStore := auth.NewAuthStore(h.db)
	authHandler := auth.NewAuthHandler(authStore)
	h.router.Handle("GET /dev", middlewares.VerifyAuthentication(Hello))
	authHandler.RegisterRoutes(h.router)

	log.Printf("server is listening on port %v", h.port)
	err = http.ListenAndServe(h.port, h.router)

	if err == nil {
		log.Fatal(err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	var k types.ContextKey = "username"
	ctx := r.Context().Value(k)
	fmt.Println(ctx)
	_, err := w.Write([]byte(ctx.(string)))
	if err != nil {
		return
	}
}
