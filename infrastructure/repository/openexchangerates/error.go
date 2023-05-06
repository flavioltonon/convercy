package openexchangerates

import "fmt"

type ResponseError struct {
	IsError     bool   `json:"error"`
	Status      int    `json:"status"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func (e ResponseError) Error() string {
	return fmt.Sprintf("%s [%d]: %s", e.Message, e.Status, e.Description)
}
