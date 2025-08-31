package domain

type Post struct {
	// 後でDBに変更する場合はjsonタグを変更
	Id      int    `json:"id"`
	Content string `json:"content"`
	Likes   int    `json:"likes"`
	UserId  int    `json:"userId"`
}
