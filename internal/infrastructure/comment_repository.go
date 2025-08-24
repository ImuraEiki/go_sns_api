package infrastructure

import (
	"encoding/json"
	"my-gin-app/internal/domain"
	"os"
)

type CommentRepository struct {
	data string
}

func NewCommentRepository() *CommentRepository {
	return &CommentRepository{data: "internal/infrastructure/data/comments.json"}
}

func (r *CommentRepository) GetAll() ([]domain.Comment, error) {
	data, err := os.ReadFile(r.data)
	if err != nil {
		return nil, err
	}
	var comments []domain.Comment
	if err := json.Unmarshal(data, &comments); err != nil {
		return nil, err
	}
	return comments, nil
}
