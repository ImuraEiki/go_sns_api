package usecase

import (
	"errors"
	"my-gin-app/internal/domain"
	"my-gin-app/internal/infrastructure"
)

type PostUsecase struct {
	// リポジトリ層のみに依存
	repo *infrastructure.PostRepository
}

func NewPostUsecase(r *infrastructure.PostRepository) *PostUsecase {
	return &PostUsecase{repo: r}
}

func (u *PostUsecase) GetAllPosts() ([]domain.Post, error) {
	return u.repo.GetAll()
}

func (u *PostUsecase) GetPostById(id int) (*domain.Post, error) {
	post, err := u.repo.GetById(id)
	if post == nil || err != nil {
		return nil, errors.New("post not found")
	}
	return u.repo.GetById(id)
}

func (u *PostUsecase) CreatePost(post *domain.Post) error {
	var newPostId int
	posts, err := u.repo.GetAll()
	if err != nil {
		return err
	}
	var maxId int
	for _, p := range posts {
		if p.Id > maxId {
			maxId = p.Id
		}
	}
	newPostId = maxId + 1

	newPost := &domain.Post{
		Id:      newPostId,
		Content: post.Content,
		Likes:   0,
		UserId:  post.UserId,
	}
	return u.repo.CreateNewPost(newPost)
}

func (u *PostUsecase) LikePost(targetPostId int) (*domain.Post, error) {
	post, err := u.repo.GetById(targetPostId)
	if post == nil || err != nil {
		return nil, errors.New("post not found")
	}
	return u.repo.LikePost(targetPostId)
}
