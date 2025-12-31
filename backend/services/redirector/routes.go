package redirector

import (
	"fmt"
	"net/http"
	"os"

	"github.com/OlyMahmudMugdho/url-shortener/services/shortener"
	"github.com/OlyMahmudMugdho/url-shortener/utils"
)

type Handler struct {
	shortenerStore shortener.LinkStore
}

func NewRedirectorHandler(store shortener.LinkStore) *Handler {
	return &Handler{
		shortenerStore: store,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	// Remove the HTTP method from the route definition
	router.HandleFunc("GET /app/{shortLink}", h.OpenLink)
}

func (h *Handler) OpenLink(w http.ResponseWriter, r *http.Request) {
	shotURl, ok := utils.ExtractParamFromUrl(r.URL.Path, "/app/")

	if !ok {
		w.WriteHeader(400)
		return
	}

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
