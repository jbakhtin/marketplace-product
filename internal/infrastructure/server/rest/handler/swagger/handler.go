package swagger

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//go:embed openapi.yaml swagger.html
var files embed.FS

func OpenAPISpec() ([]byte, error) {
	return fs.ReadFile(files, "openapi.yaml")
}

func RegisterRoutes(router chi.Router) {
	router.Get("/openapi", serveOpenAPISpec)
	router.Get("/swagger/swagger", serveSwaggerUI)
}

func serveOpenAPISpec(w http.ResponseWriter, _ *http.Request) {
	data, err := fs.ReadFile(files, "openapi.yaml")
	if err != nil {
		http.Error(w, "openapi spec not found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/yaml")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func serveSwaggerUI(w http.ResponseWriter, _ *http.Request) {
	data, err := fs.ReadFile(files, "swagger.html")
	if err != nil {
		http.Error(w, "swagger ui not found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}
