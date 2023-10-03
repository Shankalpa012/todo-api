package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"todo/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthMiddleware struct {
	userService *service.UserService
}

func NewAuthMiddleware(userService *service.UserService) *AuthMiddleware {
	return &AuthMiddleware{
		userService: userService,
	}
}

func (u AuthMiddleware) ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := getTokenFromHeader(ctx)
		if tokenHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			ctx.Abort()
			return
		}

		token, err := u.userService.ValidateToken(tokenHeader)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims["sub"])
			// Token is valid; proceed with the request
			ctx.Set("token", claims)
			ctx.Set("id", claims["sub"])

			ctx.Next()
		} else {
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}

func getTokenFromHeader(ctx *gin.Context) string {
	header := ctx.Request.Header.Get("Authorization")
	idToken := strings.TrimSpace(strings.Replace(header, "Bearer", "", 1))
	return idToken
}
