package model

// ErrorMsg ...
type ErrorMsg struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors,omitempty"`
}

// InfoMsg ...
type InfoMsg struct {
	Message string `json:"message"`
}
