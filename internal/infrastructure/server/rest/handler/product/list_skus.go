package product

import (
	"errors"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
	"net/http"
	"strconv"

	"github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/handler/response"
)

type GetSKUsListResponse struct {
	List []domain.SKU
}

func (o *Handler) List(w http.ResponseWriter, r *http.Request) {
	startSkuParam := r.URL.Query().Get("start_sku")
	if startSkuParam == "" {
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil, errors.New("empty start_sku param"))
		return
	}

	startSkuInt, err := strconv.Atoi(startSkuParam)
	if err != nil {
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil, err)
		return
	}

	countParam := r.URL.Query().Get("count")
	if countParam == "" {
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil, errors.New("empty count param"))
		return
	}

	countInt, err := strconv.Atoi(countParam)
	if err != nil {
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil, err)
		return
	}

	list, err := o.useCase.GetSKUList(r.Context(), domain.SKU(startSkuInt), countInt)

	response.WriteStandardResponse(w, r, http.StatusOK, GetSKUsListResponse{
		List: list,
	}, nil)
}
