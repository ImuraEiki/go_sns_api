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

func (r *CommentRepository) GetCommentsByPostId(postId int) ([]domain.Comment, error) {
	data, err := os.ReadFile(r.data)
	if err != nil {
		return nil, err
	}
	var comments []domain.Comment
	if err := json.Unmarshal(data, &comments); err != nil {
		return nil, err
	}
	// TODO: DBから取得する。
	var filteredComments []domain.Comment
	for _, comment := range comments {
		if comment.PostId == postId {
			filteredComments = append(filteredComments, comment)
		}
	}
	return filteredComments, nil
}

func (r *CommentRepository) CreateNewComment(comment *domain.Comment) error {
	comments, err := r.GetAll()
	if err != nil {
		return err
	}
	comments = append(comments, *comment)
	data, err := json.Marshal(comments)
	if err != nil {
		return err
	}
	if err := os.WriteFile(r.data, data, 0644); err != nil {
		return err
	}
	return nil
}
