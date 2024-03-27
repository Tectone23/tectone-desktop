package model

// Container represents a docker container
type Container struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Size   int64  `json:"size"`
	Status string `json:"status"`
}
