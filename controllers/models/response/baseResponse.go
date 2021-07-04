package response

var Success *BaseResponse

func init() {
	Success = &BaseResponse{IsSuccess: true}
}

type BaseResponse struct {
	IsSuccess 		bool
	ErrorMessage  string
}
