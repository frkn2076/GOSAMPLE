package request

type BaseRequest struct {
	Version  string `json:"version"`
	Language string `json:"language"`
}
