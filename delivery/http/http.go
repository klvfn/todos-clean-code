package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/klvfn/todos-clean-code/delivery/http/middleware"
	"github.com/klvfn/todos-clean-code/service"
)

func InitHTTP(svc *service.Service) *fiber.App {
	// Init http instance using fiber
	f := fiber.New()
	f.Use(recover.New())
	f.Use(logger.New())
	f.Use(middleware.ValidateContentType())

	// Init root router
	v1Router := f.Group("/v1")

	// Init handlers
	NewTodoHandler(f, svc, v1Router)

	return f
}
