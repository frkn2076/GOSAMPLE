package request

type AccountRequest struct {
	BaseRequest
	UserName string
	Password string
}
