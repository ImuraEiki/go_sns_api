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
	return u.repo.CreateNewUser(user)
}

func (u *UserUsecase) UpdateUser(user *domain.User) error {

	targetUser, err := u.repo.GetById(user.Id)
	if targetUser == nil || err != nil {
		return errors.New("user not found")
	}

	updatedUser := &domain.User{
		Id:      user.Id,
		Name:    user.Name,
		Email:   user.Email,
		Picture: user.Picture,
	}

	return u.repo.UpdateUser(updatedUser)
}
