package interfaces

import (
	"my-gin-app/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(u *usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.usecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	// IDを整数に変換
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.usecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
