package product

import (
	"encoding/json"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/handler/response"
)

type GetSKUsListRequest struct {
	StartSKU domain.SKU
	Count    int32
}

func (o *Handler) List(w http.ResponseWriter, r *http.Request) {
	var request GetSKUsListRequest
	_ = json.NewDecoder(r.Body).Decode(&request)

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil, err)
		return
	}

	// TODO: add logic
	// ...

	response.WriteStandardResponse(w, r, http.StatusOK, nil, nil)
}
