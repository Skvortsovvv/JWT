package middleware

import (
	"context"
	"net/http"
	"strings"

	"JWT/pkg/service"
	"JWT/pkg/store"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

type Middleware struct {
	storage  store.Store
	services *service.Service
}

func NewMiddleware(storage store.Store, services *service.Service) *Middleware {
	return &Middleware{storage: storage, services: services}
}

func (m *Middleware) UserIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			http.Error(w, "empty auth header", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			http.Error(w, "invalid auth header", http.StatusUnauthorized)
			return
		}

		if len(headerParts[1]) == 0 {
			http.Error(w, "token is empty", http.StatusUnauthorized)
			return
		}

		userId, err := m.services.ParseToken(headerParts[1])
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), userCtx, userId))

		next.ServeHTTP(w, r)
	})
}

func GetUserId(ctx context.Context) (uuid.UUID, error) {
	id := ctx.Value(userCtx)

	idStr, ok := id.(string)
	if !ok {
		return uuid.Nil, errors.New("user id is of invalid type")
	}

	return uuid.MustParse(idStr), nil
}
