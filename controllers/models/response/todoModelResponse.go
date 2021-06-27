package response

import "time"

type TodoModelResponse struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	IsCompleted bool      `json:"isCompleted"`
}
