package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jolotech/jolo-mars/internal/services/user"
	"github.com/jolotech/jolo-mars/internal/helpers"
	"github.com/jolotech/jolo-mars/types"
)

type AuthHandler struct {
	UserAuthService services.UserAuthService
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req types.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	msg, data, statusCode, err := h.UserAuthService.Register(c, req)
	if err != nil {
		helpers.ErrorResponse(c, err, msg, statusCode)
		return
	}
	helpers.SuccessResponse(c, data, msg, statusCode)
}
