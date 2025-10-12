package interfaces

import (
	"my-gin-app/internal/domain"
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

func (h *UserHandler) GetUserById(c *gin.Context) {
	idParam := c.Param("id")
	// Idを整数に変換
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user Id"})
		return
	}

	user, err := h.usecase.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser, err := h.usecase.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func (h *UserHandler) UpdateUserName(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idParam := c.Param("id")
	// Idを整数に変換
	id, err := strconv.Atoi(idParam)
	if err != nil || user.Id != id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user Id"})
		return
	}

	updatedUser, err := h.usecase.UpdateUserName(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}
