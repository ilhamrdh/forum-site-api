package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/situs-forum/internal/configs"
	"github.com/ilhamrdh/situs-forum/pkg/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(ctx *gin.Context) {
		header := ctx.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userId, username, err := jwt.ValidateToken(header, secretKey)
		log.Println("Middlewaare UserID", userId)
		log.Println("Middlewaare Username", username)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		ctx.Set("userId", userId)
		ctx.Set("username", username)
		ctx.Next()
	}
}

func AuthRefreshMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(ctx *gin.Context) {
		header := ctx.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userId, username, err := jwt.ValidateTokenWithoutExpiry(header, secretKey)
		log.Println("Middlewaare UserID", userId)
		log.Println("Middlewaare Username", username)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		ctx.Set("userId", userId)
		ctx.Set("username", username)
		ctx.Next()
	}
}
