package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"my-gin-app/internal/infrastructure"
	interfaces "my-gin-app/internal/interface"
	"my-gin-app/internal/usecase"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	// リポジトリ
	userRepo := infrastructure.NewUserRepository()
	postRepo := infrastructure.NewPostRepository()

	// ユースケース
	userUC := usecase.NewUserUsecase(userRepo)
	postUC := usecase.NewPostUsecase(postRepo)

	// ハンドラ
	userHandler := interfaces.NewUserHandler(userUC)
	postHandler := interfaces.NewPostHandler(postUC)

	// ルーティング
	// ユーザー関連のエンドポイント
	r.GET("/api/users", userHandler.GetUsers)
	r.GET("/api/users/:id", userHandler.GetUserByID)
	// 投稿関連のエンドポイント
	r.GET("/api/posts", postHandler.GetPosts)

	r.Run(":8080")
}
