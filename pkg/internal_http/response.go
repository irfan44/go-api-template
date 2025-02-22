package internal_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/irfan44/go-example/internal/dto"
)

func NewOKStatusBaseResponse() dto.BaseResponse {
	return dto.BaseResponse{
		ResponseMessage: "SUCCESS",
		ResponseCode:    http.StatusOK,
	}
}

func NewCreatedStatusBaseResponse() dto.BaseResponse {
	return dto.BaseResponse{
		ResponseMessage: "SUCCESS",
		ResponseCode:    http.StatusCreated,
	}
}

func NewAPIPath(method string, path string) string {
	return fmt.Sprintf("%s %s", method, path)
}

func SendResponse(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
