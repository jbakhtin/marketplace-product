package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	openapi "github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/openapi/v1"
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/handler/product"
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/response"
	swagger "github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/handler/swagger"
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
	productUseCase use_case.UseCaseInterface,
) (*chi.Mux, error) {
	productHandler, err := product.NewProductHandler(cfg, logger, productUseCase)
	if err != nil {
		return nil, err
	}

	openAPIValidator, err := custommiddleware.NewOpenAPIValidator()
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.URLFormat)

	router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	router.Get("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ready"))
	})

	swagger.RegisterRoutes(router)

	router.Group(func(r chi.Router) {
		r.Use(openAPIValidator)
		openapi.HandlerWithOptions(productHandler, openapi.ChiServerOptions{
			BaseRouter: r,
			ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
				response.Write(w, http.StatusBadRequest, nil, err)
			},
		})
	})

	return router, nil
}
