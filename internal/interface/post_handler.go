package interfaces

import (
    "net/http"
    "my-gin-app/internal/usecase"
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
