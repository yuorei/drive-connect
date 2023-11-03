package middleware

import (
	"context"
	"drive-connect-bff/client"

	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type authString string

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			next.ServeHTTP(w, r)
			return
		}

		bearer := "Bearer "
		auth = auth[len(bearer):]
		c := client.NewClient()
		validate, err := c.JwtValidate(auth)
		if err != nil || validate.UserId == "" {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), authString("auth"), &jwtCustomClaim{
			UserID: validate.UserId,
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func CtxValue(ctx context.Context) *jwtCustomClaim {
	raw, _ := ctx.Value(authString("auth")).(*jwtCustomClaim)
	return raw
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.Claims
}
