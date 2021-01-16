package http

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/klvfn/todos-clean-code/helper"
	miscEntity "github.com/klvfn/todos-clean-code/pkg/misc/entity"
	todoEntity "github.com/klvfn/todos-clean-code/pkg/todo/entity"
	"github.com/klvfn/todos-clean-code/pkg/todo/service"
)

// TodoHandler represent http handler for todo
type TodoHandler struct {
	TodoService service.TodoService
}

// InitTodoHandler create new instance of todo handler
func InitTodoHandler(f *fiber.App, todoService service.TodoService, rootRouter fiber.Router) {
	handler := &TodoHandler{
		TodoService: todoService,
	}

	todo := rootRouter.Group("/todo")
	todo.Get("/", handler.GetAll)
	todo.Post("/", handler.Create)
	todo.Get("/:id", handler.GetByID)
	todo.Post("/:id", handler.Delete)
	todo.Put("/:id", handler.Update)
}

// GetAll retrieve all todo
func (th TodoHandler) GetAll(c *fiber.Ctx) error {
	ctx := context.Background()
	response := miscEntity.Response{}
	response.Data = fiber.Map{}
	response.Error = fiber.Map{}

	todos, err := th.TodoService.GetAll(ctx)
	if err != nil {
		response.Error = helper.ConstructErrorResponse(http.StatusNotFound, "Failed get todos")
		response.Message = "failed"
		c.Status(http.StatusNotFound)
		return c.JSON(response)
	}

	response.Data = fiber.Map{"todos": todos}
	response.Message = "success"
	c.Status(http.StatusOK)
	return c.JSON(response)
}

// Create post a new todo
func (th TodoHandler) Create(c *fiber.Ctx) error {
	ctx := context.Background()
	payload := todoEntity.Todo{}
	response := miscEntity.Response{}
	response.Error = fiber.Map{}
	response.Data = fiber.Map{}

	err := c.BodyParser(&payload)
	if err != nil {
		response.Error = helper.ConstructErrorResponse(http.StatusBadRequest, "Invalid payload")
		response.Message = "failed"
		c.Status(http.StatusBadRequest)
		return c.JSON(response)
	}

	id, err := th.TodoService.Create(ctx, payload)
	if err != nil {
		response.Error = helper.ConstructErrorResponse(http.StatusInternalServerError, "Create todo failed")
		response.Message = "failed"
		c.Status(http.StatusInternalServerError)
		return c.JSON(response)
	}

	response.Data = fiber.Map{"confirmation": "create todo success", "id": id}
	response.Message = "success"
	return c.JSON(response)
}

// GetByID get todo by id
func (th TodoHandler) GetByID(c *fiber.Ctx) error {
	ctx := context.Background()
	response := miscEntity.Response{}
	response.Data = fiber.Map{}
	response.Error = fiber.Map{}
	id := strings.Trim(c.Params("id"), " ")
	id64, _ := strconv.ParseInt(id, 10, 64)

	todo, err := th.TodoService.GetByID(ctx, id64)
	if err != nil {
		response.Error = helper.ConstructErrorResponse(http.StatusBadRequest, "Todo not found")
		response.Message = "failed"
		c.Status(http.StatusBadRequest)
		return c.JSON(response)
	}

	response.Data = fiber.Map{"todo": todo}
	response.Message = "success"
	c.Status(http.StatusOK)
	return c.JSON(response)
}

// Delete a todo
func (th TodoHandler) Delete(c *fiber.Ctx) error {
	ctx := context.Background()
	response := miscEntity.Response{}
	response.Error = fiber.Map{}
	response.Data = fiber.Map{}
	id := strings.Trim(c.Params("id"), " ")
	id64, _ := strconv.ParseInt(id, 10, 64)

	err := th.TodoService.Delete(ctx, id64)
	if err != nil {
		response.Error = helper.ConstructErrorResponse(http.StatusInternalServerError, "Delete todo failed")
		response.Message = "failed"
		c.Status(http.StatusInternalServerError)
		return c.JSON(response)
	}

	response.Data = fiber.Map{"confirmation": "delete todo success", "id": id}
	response.Message = "success"
	return c.JSON(response)
}

// Update a todo
func (th TodoHandler) Update(c *fiber.Ctx) error {
	ctx := context.Background()
	payload := todoEntity.Todo{}
	id := strings.Trim(c.Params("id"), " ")
	id64, _ := strconv.ParseInt(id, 10, 64)
	response := miscEntity.Response{}
	response.Error = fiber.Map{}
	response.Data = fiber.Map{}

	err := c.BodyParser(&payload)
	if err != nil {
		response.Error = helper.ConstructErrorResponse(http.StatusBadRequest, "Invalid payload")
		response.Message = "failed"
		c.Status(http.StatusBadRequest)
		return c.JSON(response)
	}

	// Check first whether todo exist or not
	_, err = th.TodoService.GetByID(ctx, id64)
	if err != nil {
		response.Error = helper.ConstructErrorResponse(http.StatusBadRequest, "Todo not found, cannot be update")
		response.Message = "failed"
		c.Status(http.StatusBadRequest)
		return c.JSON(response)
	}

	payload.ID = id64
	err = th.TodoService.Update(ctx, id64, payload)
	if err != nil {
		response.Error = helper.ConstructErrorResponse(http.StatusInternalServerError, "Update todo failed")
		response.Message = "failed"
		c.Status(http.StatusInternalServerError)
		return c.JSON(response)
	}

	response.Data = fiber.Map{"confirmation": "update todo success", "id": id}
	response.Message = "success"
	return c.JSON(response)
}
