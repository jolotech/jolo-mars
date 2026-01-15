package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFoundHandler(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "The endpoint you are looking for does not exist.",
			"hint":    "Check your API route or refer to the documentation.",
			"path":    c.Request.URL.Path,
		})
	})
}

