package middleware

import (
	"errors"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"
	swagger "github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/handler/swagger"
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/response"
)

func NewOpenAPIValidator() (func(http.Handler) http.Handler, error) {
	data, err := swagger.OpenAPISpec()
	if err != nil {
		return nil, err
	}

	loader := &openapi3.Loader{IsExternalRefsAllowed: true}
	doc, err := loader.LoadFromData(data)
	if err != nil {
		return nil, err
	}

	doc.Servers = nil

	return nethttpmiddleware.OapiRequestValidatorWithOptions(doc, &nethttpmiddleware.Options{
		SilenceServersWarning: true,
		ErrorHandler: func(w http.ResponseWriter, message string, statusCode int) {
			response.Write(w, statusCode, nil, errors.New(message))
		},
	}), nil
}
