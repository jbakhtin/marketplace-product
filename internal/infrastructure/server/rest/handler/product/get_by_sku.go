package product

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/handler/response"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
)

type DeleteItemRequest struct {
	ItemSKU domain.SKU
}

func (o *Handler) Get(w http.ResponseWriter, r *http.Request) {
	var request DeleteItemRequest
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
