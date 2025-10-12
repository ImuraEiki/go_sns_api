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

func (r *UserRepository) GetById(id int) (*domain.User, error) {
	users, err := r.GetAll()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Id == id {
			return &user, nil
		}
	}
	return nil, nil
}

func (r *UserRepository) CreateNewUser(user *domain.User) (*domain.User, error) {
	var newUserId int
	users, err := r.GetAll()
	if err != nil {
		return nil, err
	}
	// TODO: DBになったらIDはオートインクリメントされるので消す
	var maxId int
	for _, u := range users {
		if u.Id > maxId {
			maxId = u.Id
		}
	}
	newUserId = maxId + 1

	newUser := &domain.User{
		Id:      newUserId,
		Name:    user.Name,
		Email:   user.Email,
		Picture: user.Picture,
	}
	users = append(users, *newUser)
	data, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}
	if err := os.WriteFile(r.data, data, 0644); err != nil {
		return nil, err
	}
	return newUser, nil
}

func (r *UserRepository) UpdateUserName(user *domain.User) (*domain.User, error) {
	users, err := r.GetAll()
	if err != nil {
		return nil, err
	}
	for i, u := range users {
		if u.Id == user.Id {
			users[i] = *user
			break
		}
	}
	data, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}
	if err := os.WriteFile(r.data, data, 0644); err != nil {
		return nil, err
	}
	return user, nil
}
