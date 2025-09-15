package product

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
	"math"
	"net/http"
	"strconv"

	"github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/handler/response"
)

type GetSKUsListRequest struct {
	StartAfterSKU string `validate:"required,numeric"`
	Count         string `validate:"required,numeric"`
}

type GetSKUsListResponse struct {
	SKUs []domain.SKU `json:"skus,omitempty"`
}

func (o *Handler) GetListSKUs(w http.ResponseWriter, r *http.Request) {
	req := GetSKUsListRequest{
		StartAfterSKU: r.URL.Query().Get("start_after_sku"),
		Count:         r.URL.Query().Get("count"),
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil,
			errors.New("validation failed: "+err.Error()))
		return
	}

	startAfterSKUInt, err := strconv.Atoi(req.StartAfterSKU)
	if err != nil {
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil, err)
		return
	}

	err = validateSKUParam(domain.SKU(startAfterSKUInt))
	if err != nil {
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil, err)
		return
	}

	countInt, err := strconv.Atoi(req.Count)
	if err != nil {
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil, err)
		return
	}

	err = validateCountParam(countInt)
	if err != nil {
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil, err)
		return
	}

	listSKUs, err := o.useCase.GetSKUList(r.Context(), domain.SKU(startAfterSKUInt), countInt)
	if err != nil {
		response.WriteStandardResponse(w, r, http.StatusInternalServerError, nil, err)
		return
	}

	response.WriteStandardResponse(w, r, http.StatusOK, GetSKUsListResponse{
		SKUs: listSKUs,
	}, nil)
}

func validateCountParam(count int) error {
	if count <= 0 {
		return errors.New("count must be positive")
	} else if count > math.MaxInt32 {
		return errors.New("count is too large")
	}

	return nil
}
