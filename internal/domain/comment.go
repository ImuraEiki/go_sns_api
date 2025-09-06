package domain

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	PostId  int    `json:"postId"`
	UserId  int    `json:"userId"`
}
