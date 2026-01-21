package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jolotech/jolo-mars/internal/services/user"
)

type AuthHandler struct {
	AuthService services.UserAuthService
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req services.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	resp, statusCode := h.AuthService.Register(c, req)
	c.JSON(statusCode, resp)
}
