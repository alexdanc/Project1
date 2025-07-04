package main

import (
	"Project1/database"
	handlers "Project1/handler"
	"Project1/internal/TaskService"
	"Project1/internal/UserService"
	tasks "Project1/internal/Web/Tasks"
	users "Project1/internal/Web/Users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	// Правильно инициализируем БД и сохраняем результат
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	repoTask := TaskService.NewTaskRepository(db)
	serviceTask := TaskService.NewTaskService(repoTask)
	handlerTask := handlers.NewTaskHandlers(serviceTask)

	repoUser := UserService.NewUserRepository(db)
	serviceUser := UserService.NewUserService(repoUser)
	handlerUser := handlers.NewUserHandler(serviceUser)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandlerTask := tasks.NewStrictHandler(handlerTask, nil)
	tasks.RegisterHandlers(e, strictHandlerTask)

	strictHandlerUser := users.NewStrictHandler(handlerUser, nil)
	users.RegisterHandlers(e, strictHandlerUser)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
