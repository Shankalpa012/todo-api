package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenHeader := getTokenFromHeader(ctx)

		if tokenHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			ctx.Abort()
			return
		}

		token, err := validateToken(tokenHeader)

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

func validateToken(tokenString string) (*jwt.Token, error) {
	// Parse and validate the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method")
		}
		return []byte(os.Getenv("SECRETE")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return token, nil
}

func getTokenFromHeader(ctx *gin.Context) string {
	header := ctx.Request.Header.Get("Authorization")
	idToken := strings.TrimSpace(strings.Replace(header, "Bearer", "", 1))
	return idToken
}
