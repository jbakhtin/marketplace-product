package product

import (
	"net/http"

	openapi "github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/openapi/v1"
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/response"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
)

func (h *Handler) ListProductSKUs(w http.ResponseWriter, r *http.Request, params openapi.ListProductSKUsParams) {
	listSKUs, err := h.useCase.GetSKUList(r.Context(), domain.SKU(params.StartAfterSku), int(params.Count))
	if err != nil {
		response.Write(w, http.StatusInternalServerError, nil, err)
		return
	}

	skus := toOpenAPISKUs(listSKUs)
	response.Write(w, http.StatusOK, openapi.GetSKUsListResponse{
		Skus: &skus,
	}, nil)
}
