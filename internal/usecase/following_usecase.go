package usecase

import (
	"errors"
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

func (u *FollowingUsecase) CreateFollowing(following *domain.Following) error {
	var newFollowingId int
	followings, err := u.repo.GetAll()
	if err != nil {
		return err
	}
	maxId := 0
	for _, cm := range followings {
		if cm.Id > maxId {
			maxId = cm.Id
		}
	}
	newFollowingId = maxId + 1

	newFollowing := &domain.Following{
		Id:         newFollowingId,
		FollowId:   following.FollowId,
		FollowedId: following.FollowedId,
	}
	return u.repo.CreateNewFollowing(newFollowing)
}

func (u *FollowingUsecase) DeleteFollowing(id int) error {
	if id == 0 {
		return errors.New("invalid id")
	}
	followings, err := u.repo.GetAll()
	if err != nil {
		return err
	}
	var newFollowings []domain.Following
	for index, following := range followings {
		if following.Id == id {
			newFollowings = append(followings[:index], followings[index+1:]...)
			break
		}
	}
	return u.repo.DeleteFollowing(newFollowings)
}
