package shortener

type Handler struct {
	store *Store
}

func NewShortenerHandler(store *Store) *Handler {
	return &Handler{
		store: store,
	}
}
