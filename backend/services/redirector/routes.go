package redirector

import "github.com/OlyMahmudMugdho/url-shortener/services/shortener"

type Handler struct {
	shortenerStore *shortener.Store
}

func NewRedirectorHandler(store *shortener.Store) *Handler {
	return &Handler{
		shortenerStore: store,
	}
}
