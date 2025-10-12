package usecase

import (
	"errors"
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

func (u *CommentUsecase) GetCommentsByPostId(postId int) ([]domain.Comment, error) {
	comments, err := u.repo.GetCommentsByPostId(postId)
	if comments == nil || err != nil {
		return []domain.Comment{}, nil
	}
	return comments, nil
}

func (u *CommentUsecase) CreateComment(comment *domain.Comment) (*domain.Comment, error) {
	if comment.Content == "" {
		return nil, errors.New(("content is required"))
	}
	if comment.PostId == 0 {
		return nil, errors.New(("postId is required"))
	}
	if comment.UserId == 0 {
		return nil, errors.New(("userId is required"))
	}
	return u.repo.CreateNewComment(comment)
}
