package response

var Success *BaseResponse

func init() {
	Success = &BaseResponse{IsSuccess: true}
}

type BaseResponse struct {
	IsSuccess 		bool   	`json:"isSuccess"`
	ErrorMessage  string 	`json:"errorMessage"`
}
