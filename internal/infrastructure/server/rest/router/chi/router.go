package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/handler/product"
	custommiddleware "github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/middleware"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/ports"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/use_case"
)

type Config interface {
	GetAppKey() string
}

func NewRouter(
	cfg Config,
	logger ports.Logger,
	cartUseCase use_case.ProductUseCase,
) (*chi.Mux, error) {
	cartHandler, err := product.NewProductHandler(cfg, logger, cartUseCase)
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.URLFormat)

	// health endpoints (no auth)
	router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	router.Get("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ready"))
	})

	// product routes with auth middleware
	authMiddleware := custommiddleware.NewAuthMiddleware(cfg)
	router.Route("/product", func(r chi.Router) {
		r.Use(authMiddleware.Auth)
		r.Get("/get", cartHandler.Get)
		r.Post("/list", cartHandler.List)
	})

	return router, nil
}
