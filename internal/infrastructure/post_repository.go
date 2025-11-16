package infrastructure

import (
	"gorm.io/gorm"

	"my-gin-app/internal/domain"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) GetAll() ([]domain.Post, error) {
	var posts []domain.Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) GetById(id int) (*domain.Post, error) {
	var post domain.Post
	if err := r.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) CreateNewPost(post *domain.Post) (*domain.Post, error) {
	if err := r.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) LikePost(targetPostId int) (*domain.Post, error) {
	var post domain.Post
	if err := r.db.First(&post, targetPostId).Error; err != nil {
		return nil, err
	}
	post.Likes++
	if err := r.db.Save(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
