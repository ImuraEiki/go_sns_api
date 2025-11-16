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

type DeleteFollowingResponse struct {
	Id      int  `json:"id"`
	Deleted bool `json:"deleted"`
}

func NewFollowingUsecase(r *infrastructure.FollowingRepository) *FollowingUsecase {
	return &FollowingUsecase{repo: r}
}

func (u *FollowingUsecase) GetAllFollowings() ([]domain.Following, error) {
	return u.repo.GetAll()
}

func (u *FollowingUsecase) GetFollowingsByUserId(userId int) ([]domain.Following, error) {
	followings, err := u.repo.GetFollowingsByUserId(userId)
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

func (u *FollowingUsecase) DeleteFollowing(id int) (*DeleteFollowingResponse, error) {
	if id == 0 {
		return nil, errors.New("invalid id")
	}
	followings, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var deleteFollowings []domain.Following
	for index, following := range followings {
		if following.Id == id {
			deleteFollowings = append(deleteFollowings, followings[index])
			break
		}
	}
	if len(deleteFollowings) > 0 {
		repoErr := u.repo.DeleteFollowing(id)
		if repoErr != nil {
			return nil, err
		}
		resp := &DeleteFollowingResponse{
			Id:      id,
			Deleted: true,
		}
		return resp, nil
	} else {
		return nil, errors.New("following id not found")
	}
}
