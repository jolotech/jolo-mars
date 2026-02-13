package middlewares

import (
	"net/http"
	"strings"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jolotech/jolo-mars/internal/helpers"
)

func UserAuthMiddleware() gin.HandlerFunc {
	secret := os.Getenv("JWT_SECRET")

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			helpers.ErrorResponse(c, nil, "Missing Authorization header", http.StatusUnauthorized)
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			helpers.ErrorResponse(c, nil, "Authorization header must start with 'Bearer '", http.StatusUnauthorized)
			c.Abort()
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		log.Println("Extracted Token:", tokenString)
		if tokenString == "" {
			helpers.ErrorResponse(c, nil, "Token not provided in Authorization header", http.StatusUnauthorized)
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil || !token.Valid {
			helpers.ErrorResponse(c, err, "Invalid or expired token", http.StatusUnauthorized)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			helpers.ErrorResponse(c, nil, "Invalid token claims", http.StatusUnauthorized)
			c.Abort()
			return
		}

		email, ok := claims["email"].(string)
		userId, ok := claims["user_id"].(string)
		if !ok || email == "" {
			helpers.ErrorResponse(c, nil, "Invalid token payload", http.StatusUnauthorized)
			c.Abort()
			return
		}

		// pass email/userId forward
		c.Set("userEmail", email)
		c.Set("userId", userId)

		c.Next()
	}
}
