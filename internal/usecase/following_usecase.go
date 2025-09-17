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

func (u *FollowingUsecase) GetFollowingsByUserId(postId int) ([]domain.Following, error) {
	followings, err := u.repo.GetFollowingsByUserId(postId)
	if followings == nil || err != nil {
		return nil, errors.New("following not found")
	}
	return followings, nil
}

func (u *FollowingUsecase) CreateFollowing(following *domain.Following) (*domain.Following, error) {
	if following.FollowUserId == 0 {
		return nil, errors.New("FollowUserId is required")
	}
	if following.FollowedUserId == 0 {
		return nil, errors.New("FollowedUserId is required")
	}
	// 重複チェック
	followings, _ := u.repo.GetAll()
	for _, val := range followings {
		if val.FollowUserId == following.FollowUserId && val.FollowedUserId == following.FollowedUserId {
			return nil, errors.New("already exists")
		}
	}
	return u.repo.CreateNewFollowing(following)
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
