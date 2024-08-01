package redirector

import (
	"fmt"
	"net/http"
	"os"
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
	// Remove the HTTP method from the route definition
	router.HandleFunc("GET /app/{shortLink}", h.OpenLink)
}

func (h *Handler) OpenLink(w http.ResponseWriter, r *http.Request) {
	url := r.RequestURI
	shotURl, _ := strings.CutPrefix(url, "/app/")

	link, err := h.shortenerStore.GetPublicLink(shotURl)

	if err != nil {
		w.Header().Add("Content-Type", "text/html")
		w.WriteHeader(404)
		data, _ := os.ReadFile("views/404.html")
		w.Write(data)
		return
	}

	if err != nil {
		w.WriteHeader(404)
		data, _ := os.ReadFile("views/404.html")
		_, err := w.Write(data)
		if err != nil {
			return
		}
		return
	}
	fmt.Println(link.FullUrl)
	http.Redirect(w, r, link.FullUrl, http.StatusSeeOther)
}
