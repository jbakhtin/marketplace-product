package product

import (
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	openapi "github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/openapi/v1"
	custommiddleware "github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/middleware"
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/response"
)

func (suite *ProductHandlerTestSuite) serveRequest(method, path string) *httptest.ResponseRecorder {
	suite.T().Helper()

	openAPIValidator, err := custommiddleware.NewOpenAPIValidator()
	suite.Require().NoError(err)

	router := chi.NewRouter()
	router.Use(middleware.URLFormat)
	router.Group(func(r chi.Router) {
		r.Use(openAPIValidator)
		openapi.HandlerWithOptions(suite.handler, openapi.ChiServerOptions{
			BaseRouter: r,
			ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
				response.Write(w, http.StatusBadRequest, nil, err)
			},
		})
	})

	req := httptest.NewRequest(method, path, nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	return rec
}
