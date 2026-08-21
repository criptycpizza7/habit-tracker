package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/criptycpizza7/habit-tracker/common/users"
	"github.com/criptycpizza7/habit-tracker/http/internal/web/helpers"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			helpers.WriteEmptyError(w, http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authorization, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			helpers.WriteEmptyError(w, http.StatusUnauthorized)
			return
		}

		token_string := parts[1]

		token, err := jwt.Parse(token_string, users.KeyFunc)

		switch {
		case token.Valid:
			raw_user_id, ok := token.Claims.(jwt.MapClaims)["user_id"].(string)
			if !ok {
				helpers.WriteEmptyError(w, http.StatusUnauthorized)
				return
			}
			user_id := users.UserId{}
			err = user_id.FromString(raw_user_id)
			if err != nil {
				helpers.WriteEmptyError(w, http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), "user_id", user_id)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		case errors.Is(err, jwt.ErrTokenMalformed):
			helpers.WriteEmptyError(w, http.StatusUnauthorized)
			return
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			helpers.WriteEmptyError(w, http.StatusUnauthorized)
			return
		case errors.Is(err, jwt.ErrTokenExpired):
			helpers.WriteEmptyError(w, http.StatusForbidden)
			return
		default:
			helpers.WriteEmptyError(w, http.StatusInternalServerError)
			return
		}
	})
}
