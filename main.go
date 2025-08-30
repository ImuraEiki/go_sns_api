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
	commentRepo := infrastructure.NewCommentRepository()
	followingRepo := infrastructure.NewFollowingRepository()

	// ユースケース
	userUC := usecase.NewUserUsecase(userRepo)
	postUC := usecase.NewPostUsecase(postRepo)
	commentUC := usecase.NewCommentUsecase(commentRepo)
	followingUC := usecase.NewFollowingUsecase(followingRepo)

	// ハンドラ
	userHandler := interfaces.NewUserHandler(userUC)
	postHandler := interfaces.NewPostHandler(postUC)
	commentHandler := interfaces.NewCommentHandler(commentUC)
	followingHandler := interfaces.NewFollowingHandler(followingUC)

	// ルーティング
	// ユーザー関連のエンドポイント
	users := r.Group("/api/users")
	{
		users.GET("", userHandler.GetUsers)
		users.GET("/:id", userHandler.GetUserByID)
		users.POST("", userHandler.CreateUser)
	}
	// 投稿関連のエンドポイント
	posts := r.Group("/api/posts")
	{
		posts.GET("", postHandler.GetPosts)
		posts.POST("", postHandler.CreatePost)
	}
	// コメント関連のエンドポイント
	comments := r.Group("/api/comments")
	{
		comments.GET("", commentHandler.GetComments)
	}
	// フォロー関連のエンドポイント
	followings := r.Group("/api/followings")
	{
		followings.GET("", followingHandler.GetFollowings)
	}
	// サーバー起動
	r.Run(":8080")
}
