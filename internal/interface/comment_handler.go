package interfaces

import (
	"my-gin-app/internal/usecase"
	"net/http"

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
