package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jolotech/jolo-mars/internal/helpers"
	"github.com/jolotech/jolo-mars/internal/services/admin"
	"github.com/jolotech/jolo-mars/internal/helpers/validations"
	"github.com/jolotech/jolo-mars/types"
)

type AdminAuthHandler struct {
	AdminAuthService *admin_services.AdminAuthService
}

func NewAdminAuthHandler(svc *admin_services.AdminAuthService) *AdminAuthHandler {
	return &AdminAuthHandler{AdminAuthService: svc}
}

func (h *AdminAuthHandler) Login(c *gin.Context) {
	var req types.AdminLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		msg := validations.HandleValidationError(err)
		helpers.ErrorResponse(c, err, msg, http.StatusBadRequest)
		return
	}

	msg, data, statusCode, err := h.AdminAuthService.Login(req)
	if err != nil {
		helpers.ErrorResponse(c, err, msg, statusCode)
		return
	}
	helpers.SuccessResponse(c, data, msg, statusCode)
}



func (h *AdminAuthHandler) ChangePassword(c *gin.Context) {
	var req types.AdminChangePasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		msg := validations.HandleValidationError(err)
		helpers.ErrorResponse(c, err, msg, http.StatusBadRequest)
		return
	}

	// Authorization: Bearer <setup_token>
	auth := c.GetHeader("Authorization")
	token := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer"))
	if token == "" {
		helpers.ErrorResponse(c, nil, "missing authorization token", http.StatusUnauthorized)
		return
	}

	msg, data, statusCode, err := h.AdminAuthService.ChangePassword(c, token, req)
	if err != nil {
		helpers.ErrorResponse(c, err, msg, statusCode)
		return
	}
	helpers.SuccessResponse(c, data, msg, statusCode)
}
