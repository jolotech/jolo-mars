package auth

import (
	// "net/http"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jolotech/jolo-mars/internal/helpers"
	"github.com/jolotech/jolo-mars/internal/helpers/validations"
	services "github.com/jolotech/jolo-mars/internal/services/user"
	"github.com/jolotech/jolo-mars/types"
)

type UserAuthHandler struct {
	UserAuthService *services.UserAuthService
}

func NewUserAuhHandler(userAuthService *services.UserAuthService) *UserAuthHandler {
	return &UserAuthHandler{
		UserAuthService: userAuthService,
	}
}


// ================= REGISER =======================

func (h *UserAuthHandler) Register(c *gin.Context) {
	var req types.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
	    msg := validations.HandleValidationError(err)
		helpers.ErrorResponse(c, err, msg, http.StatusBadRequest)
		return
	}

	msg, data, statusCode, err := h.UserAuthService.Register(c, req)
	if err != nil {
		helpers.ErrorResponse(c, err, msg, statusCode)
		return
	}
	helpers.SuccessResponse(c, data, msg, statusCode)
}


// ================== VERIFY OTP =====================

func (h *UserAuthHandler) VerifyOTP(c *gin.Context) {
	var req types.VerifyOTPRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		msg := validations.HandleValidationError(err)
		helpers.ErrorResponse(c, err, msg, http.StatusBadRequest)
		return
	}

	msg, data, statusCode, err := h.UserAuthService.VerifyOTP(req)
	if err != nil {
		helpers.ErrorResponse(c, err, msg, statusCode)
		return
	}

	helpers.SuccessResponse(c, data, msg, statusCode)
}

//============== RESEND OTP HANDLER =====================

func (h *UserAuthHandler) ResendOTP(c *gin.Context){
	var req types.ResendOTPRequest
	
	if err := c.ShouldBind(&req); err != nil {
		msg := validations.HandleValidationError(err)
		helpers.ErrorResponse(c, err, msg, http.StatusBadGateway)
		return
	}

	msg, data, statusCode, err := h.UserAuthService.ResendOTP(req)
	if err != nil {
		helpers.ErrorResponse(c, err, msg, statusCode)
		return
	}
	helpers.SuccessResponse(c, data, msg, statusCode)
}


// ============== FORGET PASSWORD =======================

func (h *UserAuthHandler) ForgetPassword(c *gin.Context) {
	var req types.ResendOTPRequest

	if err := c.ShouldBind(&req); err != nil {
		msg := validations.HandleValidationError(err)
		helpers.ErrorResponse(c, err, msg, http.StatusBadGateway)
		return
	}

	msg, data, statusCode, err := h.UserAuthService.ForgetPassword(req)
	if err != nil {
		helpers.ErrorResponse(c, err, msg, statusCode)
		return
	}
	helpers.SuccessResponse(c, data, msg, statusCode)
}

// ================ RESET PASSWORD =====================

func (h *UserAuthHandler) ResetPassword(c *gin.Context) {

	var req types.ResetPasswordSubmitRequest

	if err := c.ShouldBind(&req); err != nil {
		msg := validations.HandleValidationError(err)
		helpers.ErrorResponse(c, err, msg, http.StatusBadGateway)
	}

	msg, data, statusCode, err := h.UserAuthService.ResetPassword(req)
	if err != nil {
		helpers.ErrorResponse(c, err, msg, statusCode)
		return
	}
	helpers.SuccessResponse(c, data, msg, statusCode)
}

