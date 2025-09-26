package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func New() *validator.Validate {
	return validator.New()
}

func GetField(fe validator.FieldError, t int) string {
	var builder strings.Builder

	for i, char := range fe.Field() {
		if i > 0 && char >= 'A' && char <= 'Z' {
			if t == 1 {
				builder.WriteRune('_')
			} else {
				builder.WriteRune(' ')
			}
		}
		builder.WriteRune(char)
	}

	return strings.ToLower(builder.String())
}

func GetError(ve validator.ValidationErrors) any {
	out := make(map[string]string)
	for _, fe := range ve {
		out[GetField(fe, 1)] = GetErrorMsg(fe)
	}

	return out
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("the %s field must not be left blank.", GetField(fe, 0))
	case "eqfield":
		return "password and confirm password doesn't match"
	case "min":
		return fmt.Sprintf("the %s length should be greater than %s", GetField(fe, 0), fe.Param())
	case "max":
		return fmt.Sprintf("the %s length should be less than %s", GetField(fe, 0), fe.Param())
	case "email":
		return "the email field must be an email"
	case "e164":
		return "phone number must be in E.164 format (e.g., +62XXXXXX)."
	case "unique":
		return fmt.Sprintf("the %s is already taken. please choose a different one", GetField(fe, 0))
	default:
		return fmt.Sprintf("validation failed on the %s field", GetField(fe, 0))
	}
}
