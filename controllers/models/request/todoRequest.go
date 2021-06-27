package request

import "time"

type TodoRequest struct {
	BaseRequest
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	IsCompleted bool      `json:"isCompleted"`
}
