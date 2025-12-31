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
	if _, err := os.Stat(".env"); err == nil {
		err := gotenv.Load()
		if err != nil {
			log.Println("Error loading .env file:", err)
		}
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
		log.Fatal("Error connecting to database: ", err)
	} else {
		log.Println("Connected to database")
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
	// h.router.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve index.html for all other routes or handle redirects
	h.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// Check for static file existence first
		if _, err := os.Stat(filepath.Join("./dist/url-shortener-frontend/browser", path)); err == nil && path != "/" {
			fs.ServeHTTP(w, r)
			return
		}

		// Try to find short link if it's not root
		if len(path) > 1 {
			shortCode := path[1:]
			link, err := shortenerStore.GetPublicLink(shortCode)
			if err == nil {
				http.Redirect(w, r, link.FullUrl, http.StatusSeeOther)
				return
			}
		}

		// Retrieve index.html for SPA
		http.ServeFile(w, r, "./dist/url-shortener-frontend/browser/index.html")
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
