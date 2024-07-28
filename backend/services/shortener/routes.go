package shortener

import (
	"encoding/json"
	"github.com/OlyMahmudMugdho/url-shortener/middlewares"
	"github.com/OlyMahmudMugdho/url-shortener/models"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	store *Store
}

func NewShortenerHandler(store *Store) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.Handle("POST /add-url", middlewares.VerifyAuthentication(h.AddUrl))
}

func (h *Handler) AddUrl(w http.ResponseWriter, r *http.Request) {
	link := new(models.Link)

	err := json.NewDecoder(r.Body).Decode(&link)

	link.CreatedAt = time.Now()

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	savedLink, err := h.store.SaveLink(link)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(savedLink)
}
