package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header missing"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // "Bearer " が無かった
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			return
		}

		domain := os.Getenv("AUTH0_DOMAIN")
		audience := os.Getenv("AUTH0_AUDIENCE")

		// Auth0 の公開鍵(JWKS)を取得して検証
		jwksURL := fmt.Sprintf("%s.well-known/jwks.json", domain)
		jwks, _ := keyfunc.Get(jwksURL, keyfunc.Options{})

		token, err := jwt.Parse(tokenString, jwks.Keyfunc, jwt.WithAudience(audience), jwt.WithIssuer(domain))
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// ユーザー情報をコンテキストに保存して下流に渡す
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			c.Set("user", claims)
		}

		c.Next()
	}
}
