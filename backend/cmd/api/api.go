package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/OlyMahmudMugdho/url-shortener/middlewares"
	"github.com/OlyMahmudMugdho/url-shortener/services/redirector"
	"github.com/OlyMahmudMugdho/url-shortener/services/shortener"

	"github.com/OlyMahmudMugdho/gotenv/gotenv"
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

	err = utils.CreateTables(h.db)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("table created")

	authStore := auth.NewAuthStore(h.db)
	authHandler := auth.NewAuthHandler(authStore)
	authHandler.RegisterRoutes(h.router)

	shortenerStore := shortener.NewShortenerStore(h.db)
	shortenerHandler := shortener.NewShortenerHandler(shortenerStore)
	shortenerHandler.RegisterRoutes(h.router)

	redirectorHandler := redirector.NewRedirectorHandler(shortenerStore)
	redirectorHandler.RegisterRoutes(h.router)

	fs := http.FileServer(http.Dir("./dist/url-shortener-frontend/browser"))
	h.router.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve index.html for all other routes
	h.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(filepath.Join("./dist/url-shortener-frontend/browser", r.URL.Path)); os.IsNotExist(err) {
			http.ServeFile(w, r, "./dist/url-shortener-frontend/browser/index.html")
		} else {
			fs.ServeHTTP(w, r)
		}
	})

	//h.router.Handle("GET /dev", middlewares.VerifyAuthentication(Hello))

	log.Printf("server is listening on port %v", h.port)
	err = http.ListenAndServe(h.port, middlewares.CORS(h.router))

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
