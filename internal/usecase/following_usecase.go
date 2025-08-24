package usecase

import (
	"my-gin-app/internal/domain"
	"my-gin-app/internal/infrastructure"
)

type FollowingUsecase struct {
	// リポジトリ層のみに依存
	repo *infrastructure.FollowingRepository
}

func NewFollowingUsecase(r *infrastructure.FollowingRepository) *FollowingUsecase {
	return &FollowingUsecase{repo: r}
}

func (u *FollowingUsecase) GetAllFollowings() ([]domain.Following, error) {
	return u.repo.GetAll()
}
