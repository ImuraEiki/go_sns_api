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

func (u *CommentUsecase) CreateComment(comment *domain.Comment) error {
	var newCommentId int
	if comment.Id == 0 {
		comments, err := u.repo.GetAll()
		if err != nil {
			return err
		}
		maxId := 0
		for _, cm := range comments {
			if cm.Id > maxId {
				maxId = cm.Id
			}
		}
		newCommentId = maxId + 1
	} else {
		newCommentId = comment.Id
	}

	newComment := &domain.Comment{
		Id:      newCommentId,
		Content: comment.Content,
		PostId:  comment.PostId,
		UserId:  comment.UserId,
	}
	return u.repo.CreateNewComment(newComment)
}
