package shortener

import (
	"encoding/json"
	"github.com/OlyMahmudMugdho/url-shortener/models"
	"net/http"
)

type Handler struct {
	store *Store
}

func NewShortenerHandler(store *Store) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) AddUrl(w http.ResponseWriter, r *http.Response) {
	link := new(models.Link)

	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	savedLink, err := h.store.SaveLink(link)
	if err != nil {
		return
	}

	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(savedLink)
	if err != nil {
		return
	}
}
