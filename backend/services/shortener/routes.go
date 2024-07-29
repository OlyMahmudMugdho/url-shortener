package shortener

import (
	"encoding/json"
	"fmt"
	"github.com/OlyMahmudMugdho/url-shortener/types"
	"github.com/OlyMahmudMugdho/url-shortener/utils"
	"log"
	"net/http"
	"time"

	"github.com/OlyMahmudMugdho/url-shortener/middlewares"
	"github.com/OlyMahmudMugdho/url-shortener/models"
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

	context := r.Context()
	var userIdContext types.ContextKey = "userId"
	userId := utils.GetValueFromContext(context, userIdContext)

	link.UserId = userId

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	link.CreatedAt = time.Now()

	savedLink, err := h.store.SaveLink(link)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(savedLink)
	err = json.NewEncoder(w).Encode(savedLink)
	if err != nil {
		w.WriteHeader(500)
		return
	}
}
