package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	openapi "github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/openapi/v1"
)

func Write(w http.ResponseWriter, status int, payload any, err error) {
	w.Header().Set("Content-Type", "application/json")

	var body []byte
	var marshalErr error

	if err != nil {
		resp := NewErrorResponse(status, err)
		status = resp.Status
		body, marshalErr = json.Marshal(resp)
	} else {
		data, dataErr := json.Marshal(payload)
		if dataErr != nil {
			fmt.Println(dataErr.Error())
			return
		}
		raw := json.RawMessage(data)
		body, marshalErr = json.Marshal(openapi.SuccessResponse{
			Success: true,
			Status:  status,
			Data:    &raw,
		})
	}

	if marshalErr != nil {
		fmt.Println(marshalErr.Error())
		return
	}

	w.WriteHeader(status)
	if _, writeErr := w.Write(body); writeErr != nil {
		fmt.Println(writeErr.Error())
	}
}

func NewErrorResponse(status int, err error) openapi.ErrorResponse {
	return openapi.ErrorResponse{
		Success: false,
		Status:  status,
		Error: openapi.Error{
			Title: err.Error(),
		},
	}
}
