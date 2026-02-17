package auth

import (
	"net/http"
	// "strings"
	"log"

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



// ==================== ADMIN LOGIN =====================

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


// ==================== 2FA SETUP =====================
func (h *AdminAuthHandler) Setup2FA(c *gin.Context) {
	adminId := c.GetString("adminId")
	log.Println("adminId from token:.......", adminId)

	msg, data, statusCode, err := h.AdminAuthService.Setup2FA(adminId)
	if err != nil {
		helpers.ErrorResponse(c, err, msg, statusCode)
		return
	}
	helpers.SuccessResponse(c, data, msg, statusCode)
}


// ==================== 2FA CONFIRMATION =====================
func (h *AdminAuthHandler) Confirm2FA(c *gin.Context) {
	adminId := c.GetString("adminId")

	var req types.AdminTwoFAConfirmRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		msg := validations.HandleValidationError(err)
		helpers.ErrorResponse(c, err, msg, http.StatusBadRequest)
		return
	}

	msg, data, statusCode, err := h.AdminAuthService.Confirm2FA(adminId, req.Code)
	
	if err != nil {
		helpers.ErrorResponse(c, err, msg, statusCode)
		return
	}

	helpers.SuccessResponse(c, data, msg, statusCode)
}


// ==================== CHANGE PASSWORD =====================
func (h *AdminAuthHandler) ChangePassword(c *gin.Context) {
	var req types.AdminChangePasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		msg := validations.HandleValidationError(err)
		helpers.ErrorResponse(c, err, msg, http.StatusBadRequest)
		return
	}

	email, exists := helpers.GetAdminEmailFromContext(c)
	if !exists {
		helpers.ErrorResponse(c, nil, "UnAuthorized access", http.StatusUnauthorized)
		return
	}

	req.Email = email

	msg, data, statusCode, err := h.AdminAuthService.ChangePassword(req)
	if err != nil {
		helpers.ErrorResponse(c, err, msg, statusCode)
		return
	}
	helpers.SuccessResponse(c, data, msg, statusCode)
}

// ==================== DELETE ADMIN BY EMAIL (FOR TESTING) =====================
func (h *AdminAuthHandler) DeleteAdmin(c *gin.Context) {
	var req types.DeleteAdminRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		msg := validations.HandleValidationError(err)
		helpers.ErrorResponse(c, err, msg, http.StatusBadRequest)
		return
	}

	msg, data, statusCode, err := h.AdminAuthService.DeleteAdminByEmail(req)
	if err != nil {
		helpers.ErrorResponse(c, err, msg, statusCode)
		return
	}
	helpers.SuccessResponse(c, data, msg, statusCode)
}