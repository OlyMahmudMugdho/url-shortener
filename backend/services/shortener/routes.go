package shortener

import (
	"encoding/json"
	"fmt"
	"github.com/OlyMahmudMugdho/url-shortener/types"
	"github.com/OlyMahmudMugdho/url-shortener/utils"
	"log"
	"net/http"
	"strconv"
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
	router.Handle("GET /links", middlewares.VerifyAuthentication(h.GetAllLinks))
	router.Handle("GET /links/{urlId}", middlewares.VerifyAuthentication(h.GetLink))
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

func (h *Handler) GetAllLinks(w http.ResponseWriter, r *http.Request) {
	context := r.Context()
	userId := utils.GetUserIdFromContext(context)
	links, err := h.store.GetAllLinks(userId)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(links)
	if err != nil {
		return
	}
}

func (h *Handler) GetLink(w http.ResponseWriter, r *http.Request) {
	userIdStr, ok := utils.ExtractLinkIdFromUrl(r.URL.Path, "/links/")

	if !ok {
		w.WriteHeader(400)
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	link, err := h.store.GetLink(userId)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	err = json.NewEncoder(w).Encode(link)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}
