package domain

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email,omitempty"`
	Picture string `json:"picture,omitempty"`
}
