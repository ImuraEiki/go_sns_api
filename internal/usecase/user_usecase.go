package usecase

import (
	"errors"
	"my-gin-app/internal/domain"
	"my-gin-app/internal/infrastructure"
)

type UserUsecase struct {
	repo *infrastructure.UserRepository
}

func NewUserUsecase(r *infrastructure.UserRepository) *UserUsecase {
	return &UserUsecase{repo: r}
}

func (u *UserUsecase) GetAllUsers() ([]domain.User, error) {
	return u.repo.GetAll()
}

func (u *UserUsecase) GetUserById(id int) (*domain.User, error) {
	user, err := u.repo.GetById(id)
	if user == nil || err != nil {
		return nil, errors.New("user not found")
	}
	return u.repo.GetById(id)
}

func (u *UserUsecase) CreateUser(user *domain.User) (*domain.User, error) {
	if user.Name == "" {
		return nil, errors.New("userName is required")
	}
	if user.Email == "" {
		return nil, errors.New("userEmail is required")
	}
	// 重複チェック
	users, _ := u.repo.GetAll()
	for _, val := range users {
		if val.Name == user.Name && val.Email == user.Email {
			return nil, errors.New("already exists")
		}
	}
	return u.repo.CreateNewUser(user)
}

func (u *UserUsecase) UpdateUserName(user *domain.User) (*domain.User, error) {

	targetUser, err := u.repo.GetById(user.Id)
	if targetUser == nil || err != nil || targetUser.Email != user.Email {
		return nil, errors.New("user not found")
	}

	if user.Name == "" {
		return nil, errors.New("userName is required")
	}

	updatedUser := &domain.User{
		Id:      targetUser.Id,
		Name:    user.Name,
		Email:   targetUser.Email,
		Picture: targetUser.Picture,
	}

	return u.repo.UpdateUserName(updatedUser)
}
