package response

type TodosResponse struct {
	BaseResponse
	Todos []TodoModelResponse
}
