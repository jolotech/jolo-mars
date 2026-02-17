package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jolotech/jolo-mars/internal/helpers"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	secret := os.Getenv("ADMIN_JWT_SECRET")

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
		if tokenString == "" {
			helpers.ErrorResponse(c, nil, "Token not provided in Authorization header", http.StatusUnauthorized)
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// ✅ Ensure the token was signed with HMAC (HS256, HS384, HS512)
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenSignatureInvalid
			}
			return []byte(secret), nil
		})

		if err != nil || token == nil || !token.Valid {
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

		// Extract payload safely
		email, _ := claims["email"].(string)
		purpose, _ := claims["purpose"].(string)

		adminId, ok := claims["admin_id"].(string)
		if !ok || email == "" {
			helpers.ErrorResponse(c, nil, "Invalid token payload", http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set("adminEmail", email)
		c.Set("adminId", adminId)
		c.Set("adminPurpose", purpose)

		c.Next()
	}
}


func RequireAdminToken(allowed ...string) gin.HandlerFunc {
	allowedSet := map[string]bool{}
	for _, a := range allowed {
		allowedSet[a] = true
	}

	return func(c *gin.Context) {
		purposeVal, ok := c.Get("adminPurpose")
		if !ok {
			helpers.ErrorResponse(c, nil, "Missing token purpose", http.StatusUnauthorized)
			c.Abort()
			return
		}

		purpose, _ := purposeVal.(string)
		if purpose == "" || !allowedSet[purpose] {
			helpers.ErrorResponse(c, nil, "Invalid token purpose", http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Next()
	}
}