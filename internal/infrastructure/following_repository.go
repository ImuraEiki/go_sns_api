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

func (r *FollowingRepository) CreateNewFollowing(following *domain.Following) (*domain.Following, error) {
	var newFollowingId int
	followings, err := r.GetAll()
	if err != nil {
		return nil, err
	}
	maxId := 0
	for _, cm := range followings {
		if cm.Id > maxId {
			maxId = cm.Id
		}
	}
	newFollowingId = maxId + 1

	newFollowing := &domain.Following{
		Id:             newFollowingId,
		FollowUserId:   following.FollowUserId,
		FollowedUserId: following.FollowedUserId,
	}
	followings = append(followings, *newFollowing)
	data, err := json.Marshal(followings)
	if err != nil {
		return nil, err
	}
	if err := os.WriteFile(r.data, data, 0644); err != nil {
		return nil, err
	}
	return newFollowing, nil
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
