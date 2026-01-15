package helpers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jolotech/jolo-mars/config"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Code    int         `json:"code"`
}

// SuccessResponse sends a standardized success response
func SuccessResponse(c *gin.Context, data interface{}, message string, statusCode int) {
	if statusCode == 0 {
		statusCode = http.StatusOK
	}
	c.JSON(statusCode, Response{
		Status:  "success",
		Message: message,
		Data:    data,
		Code:    statusCode,
	})
}

// ErrorResponse sends a standardized error response with friendly messages
func ErrorResponse(c *gin.Context, err error, message string, statusCode int) {
	cfg := config.LoadConfig()

	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}

	if cfg.AppEnv == "PRODUCTION" {
		err = nil // Hide error details in production
	}

	var userFriendlyMessage = message
	var errorDetail interface{}

	if err != nil {
		errorMsg := err.Error()
		errorDetail = errorMsg

		// Detect common database errors and override message + status
		switch {
		    case strings.Contains(errorMsg, "Duplicate entry"):
			    statusCode = http.StatusConflict
			    userFriendlyMessage = extractDuplicateKeyMessage(errorMsg)
			    errorDetail = nil

		    case strings.Contains(errorMsg, "foreign key constraint fails"):
			    statusCode = http.StatusBadRequest
			    userFriendlyMessage = "Cannot perform this operation due to linked records."
			    errorDetail = nil

		    case strings.Contains(errorMsg, "record not found"):
			    statusCode = http.StatusNotFound
			    userFriendlyMessage = "The requested resource was not found."
			    errorDetail = nil

		    case strings.Contains(errorMsg, "invalid input syntax"):
			    statusCode = http.StatusBadRequest
			    userFriendlyMessage = "Invalid input format provided."
			    errorDetail = nil
		}
	}



	c.JSON(statusCode, Response{
		Status:  "error",
		Message: userFriendlyMessage,
		Error:   errorDetail,
		Code:    statusCode,
	})
}

// extractDuplicateKeyMessage parses the duplicate key error for better feedback
func extractDuplicateKeyMessage(errMsg string) string {
	start := strings.Index(errMsg, "Duplicate entry")
	if start == -1 {
		return "Duplicate entry detected"
	}

	parts := strings.Split(errMsg, "'")
	if len(parts) >= 3 {
		value := parts[1]
		return "A record with value '" + value + "' already exists."
	}
	return "Duplicate entry detected"
}
