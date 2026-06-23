package swagger

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stretchr/testify/require"
)

func TestRegisterRoutes(t *testing.T) {
	t.Parallel()

	router := chi.NewRouter()
	router.Use(middleware.URLFormat)
	RegisterRoutes(router)

	t.Run("openapi.yaml", func(t *testing.T) {
		t.Parallel()

		req := httptest.NewRequest(http.MethodGet, "/openapi.yaml", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		require.Equal(t, http.StatusOK, rec.Code)
		require.Equal(t, "application/yaml", rec.Header().Get("Content-Type"))
		require.True(t, strings.HasPrefix(rec.Body.String(), "openapi:"))
	})

	t.Run("swagger.html", func(t *testing.T) {
		t.Parallel()

		req := httptest.NewRequest(http.MethodGet, "/swagger/swagger.html", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		require.Equal(t, http.StatusOK, rec.Code)
		require.Equal(t, "text/html; charset=utf-8", rec.Header().Get("Content-Type"))
		require.Contains(t, rec.Body.String(), "swagger-ui")
		require.Contains(t, rec.Body.String(), `url: "/openapi.yaml"`)
	})
}
