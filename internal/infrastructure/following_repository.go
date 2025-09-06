package infrastructure

import (
	"encoding/json"
	"my-gin-app/internal/domain"
	"os"
)

type FollowingRepository struct {
	data string
}

func NewFollowingRepository() *FollowingRepository {
	return &FollowingRepository{data: "internal/infrastructure/data/followings.json"}
}

func (r *FollowingRepository) GetAll() ([]domain.Following, error) {
	data, err := os.ReadFile(r.data)
	if err != nil {
		return nil, err
	}
	var followings []domain.Following
	if err := json.Unmarshal(data, &followings); err != nil {
		return nil, err
	}
	return followings, nil
}

func (r *FollowingRepository) CreateNewFollowing(following *domain.Following) error {
	followings, err := r.GetAll()
	if err != nil {
		return err
	}
	followings = append(followings, *following)
	data, err := json.Marshal(followings)
	if err != nil {
		return err
	}
	if err := os.WriteFile(r.data, data, 0644); err != nil {
		return err
	}
	return nil
}

func (r *FollowingRepository) DeleteFollowing(followings []domain.Following) error {
	data, err := json.Marshal(followings)
	if err != nil {
		return err
	}
	if err := os.WriteFile(r.data, data, 0644); err != nil {
		return err
	}
	return nil
}
