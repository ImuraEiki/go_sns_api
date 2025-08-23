package domain

type Post struct {
    // 後でDBに変更する場合はjsonタグを変更
    ID      int    `json:"id"`
    Author  string `json:"author"`
    Content string `json:"content"`
}
