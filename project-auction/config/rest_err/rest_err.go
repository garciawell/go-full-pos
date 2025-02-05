package rest_err

import (
	"net/http"

	internal_error "github.com/garciawell/go-challenge-auction/internal/internal_erro"
)

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func ConvertError(internalError *internal_error.InternalError) *RestErr {
	switch internalError.Err {
	case "bad_request":
		return NewBadRequestError(internalError.Error())
	case "not_found":
		return NewNotFoundError(internalError.Error())
	default:
		return NewBadRequestError(internalError.Error())
	}
}

func NewBadRequestError(message string, causes ...Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
		Causes:  nil,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
		Causes:  nil,
	}
}
