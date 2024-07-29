package redirector

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/OlyMahmudMugdho/url-shortener/services/shortener"
)

type Handler struct {
	shortenerStore *shortener.Store
}

func NewRedirectorHandler(store *shortener.Store) *Handler {
	return &Handler{
		shortenerStore: store,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /", h.Redirect)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	url := r.RequestURI
	shotURl, _ := strings.CutPrefix(url, "/")

	link, err := h.shortenerStore.GetPublicLink(shotURl)

	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		return
	}

	fmt.Println(link.FullUrl)
	http.Redirect(w, r, link.FullUrl, http.StatusSeeOther)
}
