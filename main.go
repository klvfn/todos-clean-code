package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/klvfn/todos-clean-code/config"
	"github.com/klvfn/todos-clean-code/middleware"
	"github.com/klvfn/todos-clean-code/pkg/todo/delivery/http"
	"github.com/klvfn/todos-clean-code/pkg/todo/repository/redis"
	"github.com/klvfn/todos-clean-code/pkg/todo/service"
	"github.com/subosito/gotenv"
)

// ProjectDirectory is a root path
const ProjectDirectory = "./"

func init() {
	// Load ENV Config
	gotenv.Load(ProjectDirectory + ".env")
}

func main() {
	// Init fiber
	f := fiber.New()
	f.Use(recover.New())
	f.Use(logger.New())
	f.Use(middleware.ValidateContentType())

	// Init root router
	v1 := f.Group("/v1")

	// Init dependencies
	rdb, err := config.ConnectRedis()
	if err != nil {
		log.Fatal("Error connect to redis")
	}

	// Init context
	var contextTimeout int = 30
	if os.Getenv("CONTEXT_TIMEOUT") != "" {
		contextTimeout, err = strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
		if err != nil {
			log.Fatal("Convert context timeout failed")
		}
	}
	ctxTimeoutInSecond := time.Duration(contextTimeout) * time.Second

	// Modules todo
	redisTodoRepo := redis.NewTodoRepository(rdb)
	todoService := service.NewTodoService(redisTodoRepo, ctxTimeoutInSecond)
	http.InitTodoHandler(f, todoService, v1)

	log.Printf("Context Timeout: %ds\n", contextTimeout)
	log.Printf("[TODOS API] Listening on port %s", os.Getenv("PORT"))
	log.Fatal(f.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
