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
	"github.com/klvfn/todos-clean-code/pkg/todo/repository/mysql"
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

	// Init database
	mysqlDB, err := config.ConnectMysql()
	if err != nil {
		log.Fatal("Failed connect to mysql, Error: ", err.Error())
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
	mysqlTodoRepo := mysql.NewTodoRepository(mysqlDB)
	todoService := service.NewTodoService(mysqlTodoRepo, ctxTimeoutInSecond)
	http.InitTodoHandler(f, todoService, v1)

	log.Printf("Context Timeout: %ds\n", contextTimeout)
	log.Printf("[TODOS API] Listening on port %s", os.Getenv("PORT"))
	log.Fatal(f.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
