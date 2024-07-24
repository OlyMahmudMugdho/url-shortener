package auth

type AuthHandler struct {
	store *AuthStore
}

func NewAuthHandler(store *AuthStore) *AuthHandler {
	return &AuthHandler{
		store: store,
	}
}
