package product

import (
	"net/http"

	openapi "github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/openapi/v1"
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/response"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
	"github.com/pkg/errors"
)

func (h *Handler) GetProductBySKU(w http.ResponseWriter, r *http.Request, params openapi.GetProductBySKUParams) {
	product, err := h.useCase.GetProductBySKU(r.Context(), domain.SKU(params.Sku))
	if err != nil {
		if errors.Is(err, domain.NotFound) {
			response.Write(w, http.StatusNotFound, nil, err)
			return
		}

		h.log.Error(err.Error())
		response.Write(w, http.StatusInternalServerError, nil, err)
		return
	}

	response.Write(w, http.StatusOK, openapi.GetProductBySKUResponse{
		Product: toOpenAPIProduct(product),
	}, nil)
}
