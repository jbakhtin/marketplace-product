package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Data    any    `json:"data,omitempty"`
	Error   *Error `json:"error,omitempty"`
}

type Error struct {
	Type     string `json:"type,omitempty"`
	Title    string `json:"title"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
}

// DomainError allows mapping domain-level errors to HTTP details
type DomainError interface {
	Error() string
	Type() string
	Detail() string
	Status() int
}

func mapErrorToResponse(err error, fallbackStatus int) (int, *Error) {
	if err == nil {
		return fallbackStatus, nil
	}
	if de, ok := err.(DomainError); ok {
		return de.Status(), &Error{
			Type:   de.Type(),
			Title:  de.Error(),
			Detail: de.Detail(),
		}
	}
	return fallbackStatus, &Error{Title: err.Error()}
}

func NewSuccessResponse(status int, data any) Response {
	return Response{
		Success: true,
		Status:  status,
		Data:    data,
	}
}

func NewErrorResponse(status int, err error) Response {
	code, e := mapErrorToResponse(err, status)
	return Response{
		Success: false,
		Status:  code,
		Error:   e,
	}
}

func WriteStandardResponse(w http.ResponseWriter, r *http.Request, code int, payload any, err error) {
	var jsonResponse Response
	if err != nil {
		jsonResponse = NewErrorResponse(code, err)
		code = jsonResponse.Status
	} else {
		jsonResponse = NewSuccessResponse(code, payload)
	}

	buf, err := json.Marshal(&jsonResponse)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(buf)
	if err != nil {
		fmt.Println(err.Error())
	}
}
