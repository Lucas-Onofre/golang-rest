package main

import (
	"github.com/Lucas-Onofre/golang-rest/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/task", handlers.ListTasks)
	app.Get("/task/:id", handlers.GetTaskById)

	app.Post("/task", handlers.CreateTask)

	app.Delete("/task/:id", handlers.DeleteTaskById)
}