package response

import "time"

type TodoModelResponse struct {
	Id					uint      `json:"id"`   
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	IsCompleted bool      `json:"isCompleted"`
}
