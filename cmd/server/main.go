package main

import (
	"log"
	"todolist/internal/app/handler"
	"todolist/internal/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := database.InitDB("data/todos.db")
	defer db.Close()

	app := fiber.New()

	app.Get("/todos", handler.GetTodo(db))
	app.Post("/todos", handler.CreateTodo(db))
	app.Put("/todos", handler.UpdateTodo(db))
	app.Delete("/todos", handler.DeleteTodo(db))

	log.Fatal(app.Listen(":8080"))
}
