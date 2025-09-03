package middleware

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/handler/response"
)

type Config interface {
	GetAppKey() string
}

type CustomClaims struct {
	jwt.RegisteredClaims
}

type Middleware struct {
	cfg Config
}

func NewAuthMiddleware(cfg Config) Middleware {
	return Middleware{
		cfg: cfg,
	}
}

func (m Middleware) Auth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.WriteStandardResponse(w, r, http.StatusUnauthorized, nil, errors.New("unauthorized 0"))
			return
		}

		rawToken := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(rawToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{},
			error) {
			return []byte(m.cfg.GetAppKey()), nil
		})
		if err != nil {
			response.WriteStandardResponse(w, r, http.StatusUnauthorized, nil, errors.New(err.Error()))
			return
		}

		if !token.Valid {
			response.WriteStandardResponse(w, r, http.StatusUnauthorized, nil, errors.New("token not valid"))
			return
		}

		customClaims, ok := token.Claims.(*jwt.RegisteredClaims)
		if !ok {
			response.WriteStandardResponse(w, r, http.StatusUnauthorized, nil, errors.New("custom claims not parsed"))
			return
		}

		userID, err := strconv.Atoi(customClaims.Subject)
		if err != nil {
			response.WriteStandardResponse(w, r, http.StatusUnauthorized, nil, errors.New("token not valid"))
			return
		}
		ctx := context.WithValue(r.Context(), "user_id", userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
