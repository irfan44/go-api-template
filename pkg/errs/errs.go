package errs

import "net/http"

type MessageErr interface {
	Code() string
	StatusCode() int
	Error() string
}

type ErrorData struct {
	ResponseCode    int    `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ErrMessage      string `json:"message"`
}

func (e *ErrorData) Code() string {
	return e.ResponseMessage
}

func (e *ErrorData) StatusCode() int {
	return e.ResponseCode
}

func (e *ErrorData) Error() string {
	return e.ErrMessage
}

func NewUnauthorizedError(message string) MessageErr {
	return &ErrorData{
		ResponseMessage: "FORBIDDEN_ACCESS",
		ResponseCode:    http.StatusForbidden,
		ErrMessage:      message,
	}
}

func NewUnauthenticatedError(message string) MessageErr {
	return &ErrorData{
		ResponseMessage: "UNAUTHORIZED",
		ResponseCode:    http.StatusUnauthorized,
		ErrMessage:      message,
	}
}

func NewConflictError(message string) MessageErr {
	return &ErrorData{
		ResponseMessage: "CONFLICT",
		ResponseCode:    http.StatusConflict,
		ErrMessage:      message,
	}
}

func NewNotFoundError(message string) MessageErr {
	return &ErrorData{
		ResponseMessage: "NOT_FOUND",
		ResponseCode:    http.StatusNotFound,
		ErrMessage:      message,
	}
}

func NewBadRequest(message string) MessageErr {
	return &ErrorData{
		ResponseMessage: "BAD_REQUEST",
		ResponseCode:    http.StatusBadRequest,
		ErrMessage:      message,
	}
}

func NewInternalServerError() MessageErr {
	return &ErrorData{
		ResponseMessage: "INTERNAL_SERVER_ERROR",
		ResponseCode:    http.StatusInternalServerError,
		ErrMessage:      "Something went wrong",
	}
}

func NewUnprocessibleEntityError(message string) MessageErr {
	return &ErrorData{
		ResponseMessage: "UNPROCESSABLE_ENTITY",
		ResponseCode:    http.StatusUnprocessableEntity,
		ErrMessage:      message,
	}
}

func NewTimeOutError() MessageErr {
	return &ErrorData{
		ResponseMessage: "REQUEST_TIME_OUT",
		ResponseCode:    http.StatusRequestTimeout,
		ErrMessage:      "The request took too long to process. Please try again.",
	}
}
