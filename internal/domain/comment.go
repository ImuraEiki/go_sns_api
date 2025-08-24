package domain

type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	PostID  int    `json:"postId"`
	UserID  int    `json:"userId"`
}
