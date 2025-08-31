package usecase

import (
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

func (u *PostUsecase) CreatePost(post *domain.Post) error {
	var newPostId int
	if post.Id == 0 {
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
	} else {
		newPostId = post.Id
	}

	newPost := &domain.Post{
		Id:      newPostId,
		Content: post.Content,
		Likes:   post.Likes,
		UserId:  post.UserId,
	}
	return u.repo.CreateNewPost(newPost)
}
