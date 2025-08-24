package usecase

import (
	"my-gin-app/internal/domain"
	"my-gin-app/internal/infrastructure"
)

type CommentUsecase struct {
	// リポジトリ層のみに依存
	repo *infrastructure.CommentRepository
}

func NewCommentUsecase(r *infrastructure.CommentRepository) *CommentUsecase {
	return &CommentUsecase{repo: r}
}

func (u *CommentUsecase) GetAllComments() ([]domain.Comment, error) {
	return u.repo.GetAll()
}
