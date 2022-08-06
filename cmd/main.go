package main

import (
	"fmt"
	"log"

	"github.com/klvfn/todos-clean-code/config"
	"github.com/klvfn/todos-clean-code/delivery/http"
	"github.com/klvfn/todos-clean-code/repository"
	"github.com/klvfn/todos-clean-code/repository/mysql"
	"github.com/klvfn/todos-clean-code/service"
)

func main() {
	// Init config
	config.InitConfig()

	// Init Repo
	// Mysql
	mysqlDB, err := mysql.Connect()
	if err != nil {
		log.Fatal(err)
	}
	mysqlRepo := mysql.NewMysqlRepo(mysqlDB)

	// Init dao
	dao := repository.NewDao(mysqlRepo)

	// Init service
	svc := service.NewService(dao)

	// Init delivery
	httpInstance := http.InitHTTP(svc)

	log.Printf("[TODOS API] Listening on port %d", config.AppConfig.AppPort)
	log.Fatal(httpInstance.Listen(fmt.Sprintf(":%d", config.AppConfig.AppPort)))
}
