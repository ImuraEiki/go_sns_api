package interfaces

import (
	"my-gin-app/internal/domain"
	"my-gin-app/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPリクエストを処理する一番外側の層
type PostHandler struct {
	// ユースケース層のみに依存
	usecase *usecase.PostUsecase
}

func NewPostHandler(u *usecase.PostUsecase) *PostHandler {
	return &PostHandler{usecase: u}
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	posts, err := h.usecase.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var post struct {
		Id      *int   `json:"id"`
		Content string `json:"content"`
		Likes   int    `json:"likes"`
		UserID  int    `json:"userId"`
	}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var newPostId int
	if post.Id == nil {
		posts, err := h.usecase.GetAllPosts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		maxId := 0
		for _, p := range posts {
			if p.ID > maxId {
				maxId = p.ID
			}
		}
		newPostId = maxId + 1
	} else {
		newPostId = *post.Id
	}

	newPost := &domain.Post{
		ID:      newPostId,
		Content: post.Content,
		Likes:   post.Likes,
		UserID:  post.UserID,
	}
	if post.Id != nil {
		newPost.ID = *post.Id
	}
	if err := h.usecase.CreatePost(newPost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newPost)
}
