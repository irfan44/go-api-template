package dto

type BaseResponse struct {
	ResponseCode    int    `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}
