package validations

import (
	// "net/http"
	"strings"

	// "github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// func HandleValidationError(c *gin.Context, err error) string {
// 	errs, ok := err.(validator.ValidationErrors)
// 	if !ok {
// 		return "Validation failed"
// 	}

// 	errors := make(map[string]string)

// 	for _, e := range errs {
// 		field := strings.ToLower(e.Field())
// 		switch e.Tag() {
// 		case "required":
// 			errors[field] = field + " is required"
// 		case "min":
// 			errors[field] = field + " must be at least " + e.Param() + " characters"
// 		default:
// 			errors[field] = "invalid " + field
// 		}
// 	}

	

// 	return true
// }


func HandleValidationError(err error) string {
	errs, ok := err.(validator.ValidationErrors)
	if !ok || len(errs) == 0 {
		return "Validation failed"
	}

	e := errs[0] // 👈 FIRST validation error only
	field := strings.ToLower(e.Field())

	switch e.Tag() {
	case "required":
		return field + " is required"
	case "min":
		return field + " must be at least " + e.Param() + " characters"
	default:
		return "invalid " + field
	}
}
