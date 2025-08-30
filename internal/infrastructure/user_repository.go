package infrastructure

import (
	"encoding/json"
	"os"

	"my-gin-app/internal/domain"
)

type UserRepository struct {
	data string
}

func NewUserRepository() *UserRepository {
	return &UserRepository{data: "internal/infrastructure/data/users.json"}
}

func (r *UserRepository) GetAll() ([]domain.User, error) {
	data, err := os.ReadFile(r.data)
	if err != nil {
		return nil, err
	}
	var users []domain.User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetByID(id int) (*domain.User, error) {
	users, err := r.GetAll()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, nil
}

func (r *UserRepository) CreateNewUser(user *domain.User) error {
	users, err := r.GetAll()
	if err != nil {
		return err
	}
	users = append(users, *user)
	data, err := json.Marshal(users)
	if err != nil {
		return err
	}
	if err := os.WriteFile(r.data, data, 0644); err != nil {
		return err
	}
	return nil
}
