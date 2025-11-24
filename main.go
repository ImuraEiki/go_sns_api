package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"my-gin-app/internal/db"
	"my-gin-app/internal/domain"
	"my-gin-app/internal/infrastructure"
	interfaces "my-gin-app/internal/interface"
	"my-gin-app/internal/middleware"
	"my-gin-app/internal/usecase"
)

func main() {
	r := gin.Default()
	// DB接続
	conn := db.Connect()
	// 自動マイグレーション
	conn.AutoMigrate(&domain.User{})
	conn.AutoMigrate(&domain.Post{})
	conn.AutoMigrate(&domain.Comment{})
	conn.AutoMigrate(&domain.Following{})
	// seed値投入
	infrastructure.SeedInitialData(conn)
	// corsでAuthorizationヘッダーを許可
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// リポジトリ
	userRepo := infrastructure.NewUserRepository(conn)
	postRepo := infrastructure.NewPostRepository(conn)
	commentRepo := infrastructure.NewCommentRepository(conn)
	followingRepo := infrastructure.NewFollowingRepository(conn)

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
	auth := r.Group("/backend/users")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("", userHandler.GetUsers)
		auth.GET("/:id", userHandler.GetUserById)
		auth.POST("", userHandler.CreateUser)
		auth.PUT("/:id", userHandler.UpdateUserName)
	}
	// 投稿関連のエンドポイント
	posts := r.Group("/backend/posts")
	{
		posts.GET("", postHandler.GetPosts)
		posts.GET("/:id", postHandler.GetPostById)
		posts.POST("", postHandler.CreatePost)
		posts.PUT("/likes/:id", postHandler.LikePost)
	}
	// コメント関連のエンドポイント
	comments := r.Group("/backend/comments")
	{
		comments.GET("", commentHandler.GetComments)
		comments.GET("/posts/:id", commentHandler.GetCommentsByPostId)
		comments.POST("", commentHandler.CreateComment)
	}
	// フォロー関連のエンドポイント
	followings := r.Group("/backend/followings")
	{
		followings.GET("", followingHandler.GetFollowings)
		followings.GET("/user/:id", followingHandler.GetFollowingsByUserId)
		followings.POST("", followingHandler.CreateFollowing)
		followings.DELETE("", followingHandler.DeleteFollowing)
	}

	health := r.Group("/backend/health")
	{
		health.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "OK",
			})
		})
	}
	// サーバー起動
	r.Run(":8080")
}
