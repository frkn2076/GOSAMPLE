package response

type AccountResponse struct {
	BaseResponse
	Token string `json:"token"`
}
