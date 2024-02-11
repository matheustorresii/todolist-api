package main

import (
	"log"
	"todolist/src/database"
	"todolist/src/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db := database.InitDB("src/database/todos.db")
	defer db.Close()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,DELETE",
	}))

	app.Get("/todos", handler.GetTodo(db))
	app.Post("/todos", handler.CreateTodo(db))
	app.Put("/todos", handler.UpdateTodo(db))
	app.Delete("/todos", handler.DeleteTodo(db))

	log.Fatal(app.Listen(":8080"))
}
