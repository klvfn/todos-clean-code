package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/klvfn/todos-clean-code/entity"
	"github.com/klvfn/todos-clean-code/helper"
)

// ValidateContentType check content type should by only application/json
func ValidateContentType() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if c.Method() != "GET" && string(c.Request().Header.ContentType()) != "application/json" {
			c.Status(http.StatusBadRequest)
			errResp := helper.ConstructErrorResponse(http.StatusBadRequest, "Invalid content type")
			response := entity.Response{}
			response.Data = fiber.Map{}
			response.Error = errResp
			response.Message = "failed"
			return c.JSON(response)
		}
		return c.Next()
	}
}
