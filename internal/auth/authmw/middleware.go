package authmw

import (
	"context"
	"errors"
	"github.com/aridae/gophermart-diploma/internal/auth/authctx"
	"github.com/aridae/gophermart-diploma/internal/model"
	"net/http"
	"slices"
)

var (
	unauthorizedError = errors.New("unauthorized")
)

type jwtService interface {
	ParseToken(ctx context.Context, tokenString string) (model.User, error)
}

func AuthenticateMiddleware(
	jwtService jwtService,
	whitelist []string,
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if slices.Contains(whitelist, r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}

			token := r.Header.Get("Authorization")
			if token == "" {
				http.Error(w, unauthorizedError.Error(), http.StatusUnauthorized)
				return
			}

			ctx := r.Context()

			user, err := jwtService.ParseToken(ctx, token)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			r = r.WithContext(authctx.ContextWithUser(ctx, user))

			next.ServeHTTP(w, r)
		})
	}
}
