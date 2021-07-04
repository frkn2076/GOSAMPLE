package response

import "time"

type TodoModelResponse struct {
	Id					uint
	Name        string
	Description string
	Deadline    time.Time
	IsCompleted bool
}
