package request

type AccountRequest struct {
	BaseRequest
	UserName string `json:"username"`
	Password string `json:"password"`
}
