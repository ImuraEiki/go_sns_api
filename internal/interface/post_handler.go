package interfaces

import (
	"my-gin-app/internal/domain"
	"my-gin-app/internal/usecase"
	"net/http"
	"strconv"

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

func (h *PostHandler) GetPostById(c *gin.Context) {
	idParam := c.Param("id")
	// Idを整数に変換
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post Id"})
		return
	}

	post, err := h.usecase.GetPostById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var post domain.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPost, err := h.usecase.CreatePost(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newPost)
}

func (h *PostHandler) LikePost(c *gin.Context) {
	idParam := c.Param("id")
	// Idを整数に変換
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post Id"})
		return
	}

	post, err := h.usecase.LikePost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}
