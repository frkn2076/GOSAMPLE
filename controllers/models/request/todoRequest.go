package request

import "time"

type TodoRequest struct {
	BaseRequest
	Id					uint
	Name        string
	Description string
	Deadline    time.Time
	IsCompleted bool
}
