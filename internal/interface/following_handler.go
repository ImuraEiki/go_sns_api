package interfaces

import (
	"my-gin-app/internal/domain"
	"my-gin-app/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FollowingHandler struct {
	usecase *usecase.FollowingUsecase
}

func NewFollowingHandler(u *usecase.FollowingUsecase) *FollowingHandler {
	return &FollowingHandler{usecase: u}
}

func (h *FollowingHandler) GetFollowings(c *gin.Context) {
	followings, err := h.usecase.GetAllFollowings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, followings)
}

func (h *FollowingHandler) GetFollowingsByUserId(c *gin.Context) {
	idParam := c.Param("id")
	// Idを整数に変換
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user Id"})
		return
	}

	followings, err := h.usecase.GetFollowingsByUserId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, followings)
}

func (h *FollowingHandler) CreateFollowing(c *gin.Context) {
	var following domain.Following
	if err := c.ShouldBindJSON(&following); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newFollowing, err := h.usecase.CreateFollowing(&following)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newFollowing)
}

func (h *FollowingHandler) DeleteFollowing(c *gin.Context) {
	var following domain.Following
	if err := c.ShouldBindJSON(&following); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resp, err := h.usecase.DeleteFollowing(following.Id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, resp)
	}
}
