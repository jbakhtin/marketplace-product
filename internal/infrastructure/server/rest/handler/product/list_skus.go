package product

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/handler/response"
)

type ClearRequest struct{}

func (o *Handler) List(w http.ResponseWriter, r *http.Request) {
	var request ClearRequest
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
