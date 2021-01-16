package helper

import (
	"net/http"

	"github.com/klvfn/todos-clean-code/pkg/misc/entity"
)

// ConstructErrorResponse produce a new error response
func ConstructErrorResponse(status int, message string) entity.ErrorResponse {
	return entity.ErrorResponse{
		Code:        status,
		Description: http.StatusText(status),
		Message:     message,
	}
}
