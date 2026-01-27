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

func (h *UserAuthHandler) VerifyOTP(c *gin.Context) {
	var req types.VerifyOTPRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		msg := validations.HandleValidationError(err)
		helpers.ErrorResponse(c, err, msg, http.StatusBadRequest)
		return
	}

	msg, data, statusCode, err := h.UserAuthService.VerifyOTP(c, req)
	if err != nil {
		helpers.ErrorResponse(c, err, msg, statusCode)
		return
	}

	helpers.SuccessResponse(c, data, msg, statusCode)
}

