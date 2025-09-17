package interfaces

import (
	"my-gin-app/internal/domain"
	"my-gin-app/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	usecase *usecase.CommentUsecase
}

func NewCommentHandler(u *usecase.CommentUsecase) *CommentHandler {
	return &CommentHandler{usecase: u}
}

func (h *CommentHandler) GetComments(c *gin.Context) {
	comments, err := h.usecase.GetAllComments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) GetCommentsByPostId(c *gin.Context) {
	idParam := c.Param("id")
	// Idを整数に変換
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post Id"})
		return
	}

	comments, err := h.usecase.GetCommentsByPostId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	var comment domain.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var newComment *domain.Comment
	newComment, err := h.usecase.CreateComment(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newComment)
}
