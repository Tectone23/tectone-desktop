package model

type Feedback struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	FilePath string `json:"filepath"`
	Message  string `json:"message"`
}
