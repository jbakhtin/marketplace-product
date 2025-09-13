package product

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"

	"github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/handler/response"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
	"github.com/pkg/errors"
)

type GetProductRequest struct {
	SKU string `validate:"required,numeric,min=0"`
}

type GetProductBySKUResponse struct {
	Product domain.Product `json:"product"`
}

func (o *Handler) Get(w http.ResponseWriter, r *http.Request) {
	req := GetProductRequest{
		SKU: r.URL.Query().Get("sku"),
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil,
			errors.New("validation failed: "+err.Error()))
		return
	}

	skuInt, err := strconv.Atoi(req.SKU)
	if err != nil {
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil, err)
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
