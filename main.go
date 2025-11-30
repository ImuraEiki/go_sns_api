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
	users := r.Group("/backend/users")
	auth_users := r.Group("/backend/users")
	auth_users.Use(middleware.AuthMiddleware())
	{
		users.GET("", userHandler.GetUsers)
		users.GET("/:id", userHandler.GetUserById)
		auth_users.POST("", userHandler.CreateUser)
		auth_users.PUT("/:id", userHandler.UpdateUserName)
	}
	// 投稿関連のエンドポイント
	posts := r.Group("/backend/posts")
	auth_posts := r.Group("/backend/posts")
	auth_posts.Use(middleware.AuthMiddleware())
	{
		posts.GET("", postHandler.GetPosts)
		posts.GET("/:id", postHandler.GetPostById)
		auth_posts.POST("", postHandler.CreatePost)
		posts.PUT("/likes/:id", postHandler.LikePost)
	}
	// コメント関連のエンドポイント
	comments := r.Group("/backend/comments")
	auth_comments := r.Group("/backend/comments")
	auth_comments.Use(middleware.AuthMiddleware())
	{
		comments.GET("", commentHandler.GetComments)
		comments.GET("/posts/:id", commentHandler.GetCommentsByPostId)
		auth_comments.POST("", commentHandler.CreateComment)
	}
	// フォロー関連のエンドポイント
	followings := r.Group("/backend/followings")
	auth_followings := r.Group("/backend/followings")
	auth_followings.Use(middleware.AuthMiddleware())
	{
		followings.GET("", followingHandler.GetFollowings)
		followings.GET("/user/:id", followingHandler.GetFollowingsByUserId)
		auth_followings.POST("", followingHandler.CreateFollowing)
		auth_followings.DELETE("", followingHandler.DeleteFollowing)
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
