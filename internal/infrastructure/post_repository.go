package infrastructure

import (
	"encoding/json"
	"os"

	"my-gin-app/internal/domain"
)

type PostRepository struct {
	data string
}

func NewPostRepository() *PostRepository {
	return &PostRepository{data: "internal/infrastructure/data/posts.json"}
}

func (r *PostRepository) GetAll() ([]domain.Post, error) {
	// データをファイルから読み込み、ドメインモデルに変換する
	// 後でDBに変更する場合は、ここを変更
	data, err := os.ReadFile(r.data)
	if err != nil {
		return nil, err
	}
	var posts []domain.Post
	if err := json.Unmarshal(data, &posts); err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) CreateNewPost(post *domain.Post) error {
	// データをファイルに保存する
	posts, err := r.GetAll()
	if err != nil {
		return err
	}
	posts = append(posts, *post)
	data, err := json.Marshal(posts)
	if err != nil {
		return err
	}
	if err := os.WriteFile(r.data, data, 0644); err != nil {
		return err
	}
	return nil
}
