package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValidate(c *gin.Context, data any) map[string]string {
	if err := c.ShouldBind(data); err != nil {
		errors := parseValidationErrors(err)
		return errors
	}

	return nil
}

func parseValidationErrors(err error) map[string]string {
	errors := make(map[string]string)
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range ve {
			field := fieldErr.Field()
			tag := fieldErr.Tag()
			errors[field] = translateError(tag)
		}
	} else {
		errors["unknown"] = err.Error()
	}

	return errors
}

func translateError(tag string) string {
	switch tag {
	case "required":
		return "This field is required."
	case "email":
		return "Not a valid email address."
	default:
		return "Not a valid value."
	}
}
