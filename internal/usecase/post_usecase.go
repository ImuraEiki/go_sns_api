package usecase

import (
    "my-gin-app/internal/domain"
    "my-gin-app/internal/infrastructure"
)

type PostUsecase struct {
	// リポジトリ層のみに依存
    repo *infrastructure.PostRepository
}

func NewPostUsecase(r *infrastructure.PostRepository) *PostUsecase {
    return &PostUsecase{repo: r}
}

func (u *PostUsecase) GetAllPosts() ([]domain.Post, error) {
    return u.repo.GetAll()
}
