package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/OlyMahmudMugdho/url-shortener/types"
	"github.com/OlyMahmudMugdho/url-shortener/utils"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyAuthentication(handler http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookies := r.Cookies()
		var token = utils.ExtractToken("token", cookies)

		if len(token) == 0 {
			w.WriteHeader(403)
			return
		}

		validToken, err := utils.ValidateToken(token)

		if err != nil {
			err := json.NewEncoder(w).Encode(map[string]any{
				"error":   true,
				"message": "invalid token",
			})
			if err != nil {
				return
			}
		}
		claims := validToken.Claims.(jwt.MapClaims)
		username := claims["username"].(string)
		userId := claims["userId"].(string)

		var usernameContext types.ContextKey = "username"
		var userIdContext types.ContextKey = "userId"

		ctx := r.Context()
		ctx = context.WithValue(ctx, usernameContext, username)
		ctx = context.WithValue(ctx, userIdContext, userId)
		r = r.WithContext(ctx)

		handler.ServeHTTP(w, r)

	})
}

/*
	// unused old code
	decoded, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return t.Claims, nil
	})

	claims := decoded.Claims.(jwt.MapClaims)
	fmt.Println(claims["username"])

	fmt.Println(c)
*/
