package infrastructure

import (
	"gorm.io/gorm"

	"my-gin-app/internal/domain"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) GetAll() ([]domain.Comment, error) {
	var comments []domain.Comment
	if err := r.db.Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *CommentRepository) GetCommentsByPostId(postId int) ([]domain.Comment, error) {
	var comments []domain.Comment
	if err := r.db.Where("post_id = ?", postId).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *CommentRepository) CreateNewComment(comment *domain.Comment) (*domain.Comment, error) {
	if err := r.db.Create(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}
