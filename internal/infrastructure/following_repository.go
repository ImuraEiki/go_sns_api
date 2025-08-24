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
