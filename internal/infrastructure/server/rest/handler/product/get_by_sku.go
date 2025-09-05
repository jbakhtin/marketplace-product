package product

import (
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/handler/response"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type GetProductBySKUResponse struct {
	Product domain.Product `json:"product"`
}

func (o *Handler) Get(w http.ResponseWriter, r *http.Request) {
	skuParam := r.URL.Query().Get("sku")
	if skuParam == "" {
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil, errors.New("empty sku param"))
		return
	}

	skuInt, err := strconv.Atoi(skuParam)
	if err != nil {
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil, err)
		return
	}

	product, err := o.useCase.GetProductBySKU(r.Context(), domain.SKU(skuInt))
	if err != nil {
		o.log.Error(err.Error())
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil, err)
		return
	}

	response.WriteStandardResponse(w, r, http.StatusOK, GetProductBySKUResponse{
		Product: product,
	}, nil)
}
