package helper

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/klvfn/todos-clean-code/pkg/misc/entity"
)

// GenerateUniqueID generate unique id using uuid
func GenerateUniqueID() string {
	id := uuid.New()
	return id.String()
}

// ConstructErrorResponse produce a new error response
func ConstructErrorResponse(status int, message string) entity.ErrorResponse {
	return entity.ErrorResponse{
		Code:        status,
		Description: http.StatusText(status),
		Message:     message,
	}
}
