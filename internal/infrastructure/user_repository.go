package infrastructure

import (
	"gorm.io/gorm"

	"my-gin-app/internal/domain"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetById(id int) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateNewUser(user *domain.User) (*domain.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUserName(targetUser *domain.User) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, targetUser.Id).Error; err != nil {
		return nil, err
	}
	user.Name = targetUser.Name
	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
