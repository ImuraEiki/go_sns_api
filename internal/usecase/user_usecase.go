package usecase

import (
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

func (u *UserUsecase) GetUserByID(id int) (*domain.User, error) {
	return u.repo.GetByID(id)
}

func (u *UserUsecase) CreateUser(user *domain.User) error {
	return u.repo.CreateNewUser(user)
}
