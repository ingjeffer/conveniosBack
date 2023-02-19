package model

type Usuario struct {
	Id    string `json:"id"`
	Email string `json:"email,omitempty"`
}
