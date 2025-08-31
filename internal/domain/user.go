package domain

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email,omitempty"`
	Picture string `json:"picture,omitempty"`
}
