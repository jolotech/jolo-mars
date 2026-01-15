package middlewares

import (
	"time"

	gincors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS returns a Gin middleware using gin-contrib/cors with a sensible config.
// You can customize AllowOrigins here or read them from env vars.

// AllowOrigins: []string{
//     "http://localhost:3000",
//     "https://4ac28c34f89e.ngrok-free.app",
// },

func CORS() gin.HandlerFunc {
	return gincors.New(gincors.Config{
		AllowOrigins:     []string{"*"}, // replace with specific domains if needed
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// Optional sensible defaults:
		MaxAge: 12 * time.Hour,
	})
}
