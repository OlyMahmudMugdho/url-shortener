package shortener

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/OlyMahmudMugdho/url-shortener/types"
	"github.com/OlyMahmudMugdho/url-shortener/utils"

	"github.com/OlyMahmudMugdho/url-shortener/middlewares"
	"github.com/OlyMahmudMugdho/url-shortener/models"
)

type LinkStore interface {
	SaveLink(link *models.Link) (*models.Link, error)
	GetAllLinks(userId string) ([]models.Link, error)
	GetLink(urlId int) (*models.Link, error)
	GetPublicLink(shortUrl string) (*models.Link, error)
	UpdateLink(link *models.Link) (*models.Link, error)
	DeleteLink(urlId int, userId int) error
}

type Handler struct {
	store LinkStore
}

func NewShortenerHandler(store LinkStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.Handle("POST /add-url", middlewares.VerifyAuthentication(h.AddUrl))
	router.Handle("PUT /update-url", middlewares.VerifyAuthentication(h.UpdateLink))
	router.Handle("GET /links", middlewares.VerifyAuthentication(h.GetAllLinks))
	router.Handle("GET /links/{urlId}", middlewares.VerifyAuthentication(h.GetLink))
	router.Handle("DELETE /links/{urlId}", middlewares.VerifyAuthentication(h.DeleteLink))
}

func (h *Handler) AddUrl(w http.ResponseWriter, r *http.Request) {
	link := new(models.Link)

	err := json.NewDecoder(r.Body).Decode(&link)

	context := r.Context()
	var userIdContext types.ContextKey = "userId"
	userId := utils.GetValueFromContext(context, userIdContext)

	userIdInt, _ := strconv.Atoi(userId)
	link.UserId = userIdInt

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	if link.ShortUrl == "" {
		link.ShortUrl = utils.GenerateShortUrl(link.FullUrl)
	}

	link.CreatedAt = time.Now()

	savedLink, err := h.store.SaveLink(link)

	if err != nil {
		message := utils.DbErrorMessage(err, "url")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]any{
			"error":   true,
			"message": message,
		})
		return
	}

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
	userIdStr, ok := utils.ExtractParamFromUrl(r.URL.Path, "/links/")

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

func (h *Handler) UpdateLink(w http.ResponseWriter, r *http.Request) {
	link := new(models.Link)
	err := json.NewDecoder(r.Body).Decode(&link)
	//link.UpdatedAt = sql.NullTime{Time: time.Now().Local()}
	//fmt.Println(link.UpdatedAt.Time)
	//fmt.Println(link.CreatedAt)
	//return

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&types.Error{
			Error:   true,
			Message: "invalid json request",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Extract User ID from context (as string) and convert to int
	context := r.Context()
	var userIdContext types.ContextKey = "userId"
	userIdStr := utils.GetValueFromContext(context, userIdContext)
	userIdInt, _ := strconv.Atoi(userIdStr) // Add error handling if robust, but assuming middleware validated it

	link.UserId = userIdInt // Set the integer UserId

	link, err = h.store.UpdateLink(link)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"ok":      true,
		"message": "updated",
		"data":    &link,
	})
}

func (h *Handler) DeleteLink(w http.ResponseWriter, r *http.Request) {
	context := r.Context()
	userId := utils.GetUserIdFromContext(context)

	urlIdStr, ok := utils.ExtractParamFromUrl(r.URL.Path, "/links/")
	if !ok {
		w.WriteHeader(400)
		return
	}

	urlId, err := strconv.Atoi(urlIdStr)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	userIdInt, _ := strconv.Atoi(userId)
	err = h.store.DeleteLink(urlId, userIdInt)

	if err != nil {
		log.Println(err)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"error":   true,
			"message": "something went wrong",
		})
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]any{
		"ok":      true,
		"message": "deleted",
	})

}
