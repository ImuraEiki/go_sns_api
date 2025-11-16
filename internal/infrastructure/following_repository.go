package infrastructure

import (
	"gorm.io/gorm"

	"my-gin-app/internal/domain"
)

type FollowingRepository struct {
	db *gorm.DB
}

func NewFollowingRepository(db *gorm.DB) *FollowingRepository {
	return &FollowingRepository{db: db}
}

func (r *FollowingRepository) GetAll() ([]domain.Following, error) {
	var followings []domain.Following
	if err := r.db.Find(&followings).Error; err != nil {
		return nil, err
	}
	return followings, nil
}

func (r *FollowingRepository) GetFollowingsByUserId(userId int) ([]domain.Following, error) {
	var followings []domain.Following
	if err := r.db.Where("followed_user_id = ? OR follow_user_id = ?", userId, userId).Find(&followings).Error; err != nil {
		return nil, err
	}
	return followings, nil
}

func (r *FollowingRepository) CreateNewFollowing(following *domain.Following) (*domain.Following, error) {
	if err := r.db.Create(following).Error; err != nil {
		return nil, err
	}
	return following, nil
}

func (r *FollowingRepository) DeleteFollowing(id int) error {
	var followings []domain.Following
	if err := r.db.Delete(&followings, id).Error; err != nil {
		return err
	}
	return nil
}
