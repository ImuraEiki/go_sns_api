package domain

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Likes   int    `json:"likes"`
	UserId  int    `json:"userId"`
}
