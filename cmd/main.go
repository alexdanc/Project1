package main

import (
	"Project1/database"
	handlers "Project1/handler"
	"Project1/internal/TaskService"
	tasks "Project1/internal/Web/Tasks"
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

	repo := TaskService.NewRepository(db)
	service := TaskService.NewTaskService(repo)
	handler := handlers.NewTaskHandlers(service)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
