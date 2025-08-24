package interfaces

import (
	"my-gin-app/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FollowingHandler struct {
	usecase *usecase.FollowingUsecase
}

func NewFollowingHandler(u *usecase.FollowingUsecase) *FollowingHandler {
	return &FollowingHandler{usecase: u}
}

func (h *FollowingHandler) GetFollowings(c *gin.Context) {
	comments, err := h.usecase.GetAllFollowings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}
