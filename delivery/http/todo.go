package http

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/klvfn/todos-clean-code/entity"
	"github.com/klvfn/todos-clean-code/helper"
	"github.com/klvfn/todos-clean-code/service"
)

// TodoHandler represent http handler for todo
type TodoHandler struct {
	service *service.Service
}

// NewTodoHandler create new instance of todo handler
func NewTodoHandler(f *fiber.App, svc *service.Service, rootRouter fiber.Router) {
	handler := &TodoHandler{
		service: svc,
	}
	todo := rootRouter.Group("/todo")
	todo.Get("/", handler.GetAll)
	todo.Post("/", handler.Create)
	todo.Get("/:id", handler.GetByID)
	todo.Post("/:id", handler.Delete)
	todo.Put("/:id", handler.Update)
}

// GetAll retrieve all todo
func (h TodoHandler) GetAll(c *fiber.Ctx) error {
	response := entity.Response{}
	response.Data = fiber.Map{}

	todos, err := h.service.Todo.GetAll()
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
func (h TodoHandler) Create(c *fiber.Ctx) error {
	payload := entity.Todo{}
	response := entity.Response{}
	response.Data = fiber.Map{}

	err := c.BodyParser(&payload)
	if err != nil {
		response.Error = helper.ConstructErrorResponse(http.StatusBadRequest, "Invalid payload")
		response.Message = "failed"
		c.Status(http.StatusBadRequest)
		return c.JSON(response)
	}

	id, err := h.service.Todo.Create(payload)
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
func (h TodoHandler) GetByID(c *fiber.Ctx) error {
	response := entity.Response{}
	response.Data = fiber.Map{}
	id := strings.Trim(c.Params("id"), " ")
	id64, _ := strconv.ParseInt(id, 10, 64)

	todo, err := h.service.Todo.GetByID(id64)
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
func (h TodoHandler) Delete(c *fiber.Ctx) error {
	response := entity.Response{}
	response.Data = fiber.Map{}
	id := strings.Trim(c.Params("id"), " ")
	id64, _ := strconv.ParseInt(id, 10, 64)

	err := h.service.Todo.Delete(id64)
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
func (h TodoHandler) Update(c *fiber.Ctx) error {
	payload := entity.Todo{}
	id := strings.Trim(c.Params("id"), " ")
	id64, _ := strconv.ParseInt(id, 10, 64)
	response := entity.Response{}
	response.Data = fiber.Map{}

	err := c.BodyParser(&payload)
	if err != nil {
		response.Error = helper.ConstructErrorResponse(http.StatusBadRequest, "Invalid payload")
		response.Message = "failed"
		c.Status(http.StatusBadRequest)
		return c.JSON(response)
	}

	// Check first whether todo exist or not
	_, err = h.service.Todo.GetByID(id64)
	if err != nil {
		response.Error = helper.ConstructErrorResponse(http.StatusBadRequest, "Todo not found, cannot be update")
		response.Message = "failed"
		c.Status(http.StatusBadRequest)
		return c.JSON(response)
	}

	payload.ID = id64
	err = h.service.Todo.Update(id64, payload)
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
