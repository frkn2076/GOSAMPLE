package response

var Success *BaseResponse

func init() {
	Success = &BaseResponse{IsSuccess: true}
}

type BaseResponse struct {
	IsSuccess 		bool   `json:"isSuccess"`
	ErrorCode   	string `json:"errorCode"`
	ErrorMessage  string `json:"errorMessage"`
}
