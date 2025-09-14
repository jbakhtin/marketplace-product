package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/handler/product"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/ports"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/use_case"
)

type Config interface {
	GetAppKey() string
}

func NewRouter(
	cfg Config,
	logger ports.Logger,
	productUseCase use_case.UseCaseInterface,
) (*chi.Mux, error) {
	productHandler, err := product.NewProductHandler(cfg, logger, productUseCase)
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
	//authMiddleware := custommiddleware.NewAuthMiddleware(cfg)
	router.Route("/products", func(r chi.Router) {
		//r.Use(authMiddleware.Auth)
		r.Get("/get", productHandler.Get)
		r.Get("/list", productHandler.GetListSKUs)
	})

	return router, nil
}
