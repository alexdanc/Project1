package main

import (
	"Project1/db"
	"Project1/handler"
	"Project1/internal/TaskService"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

func main() {

	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	e := echo.New()

	taskRepo := TaskService.NewRepository(database)
	taskServ := TaskService.NewTaskService(taskRepo)
	taskHand := handlers.NewRequestBodyHandlers(taskServ)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		log.Printf("HTTP Error: %d (%s)", code, err.Error())
		c.JSON(code, map[string]string{
			"error": err.Error(),
		})
	}

	e.GET("/task", taskHand.GetHandler)
	e.POST("/task", taskHand.PostHandler)
	e.PATCH("/task/:id", taskHand.PatchHandler)
	e.DELETE("/task/:id", taskHand.DeleteHandler)

	e.Logger.Fatal(e.Start(":8080"))

}
