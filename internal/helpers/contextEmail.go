package helpers

import (
	// "net/http"
	"github.com/gin-gonic/gin"
)

// GetUserEmailFromContext extracts the partnerEmail from Gin context
func GetEmailFromContext(c *gin.Context) (string, bool) {
	emailVal, exists := c.Get("userEmail")
	if !exists {
		return "", false
	}

	email, ok := emailVal.(string)
	if !ok {
		return "", false
	}
	
	return email, true
}


func GetAdminEmailFromContext(c *gin.Context) (string, bool) {
	emailVal, exists := c.Get("adminEmail")
	if !exists {
		return "", false
	}

	email, ok := emailVal.(string)
	if !ok {
		return "", false
	}
	
	return email, true
}
