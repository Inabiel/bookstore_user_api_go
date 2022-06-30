package errors

import "net/http"

type RestfulError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `code:"error"`
}

func NewBadRequestError(message string) *RestfulError {
	return &RestfulError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}
