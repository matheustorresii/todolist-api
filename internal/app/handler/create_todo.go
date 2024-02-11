package handler

import (
	"database/sql"
	"todolist/internal/app/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateTodo(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var todo model.Todo
		if err := c.BodyParser(&todo); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		if err := insertTodoInDatabase(db, &todo); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.Status(fiber.StatusCreated).SendString("Created")
	}
}

func insertTodoInDatabase(db *sql.DB, todo *model.Todo) error {
	todo.ID = uuid.NewString()
	smtm, err := db.Prepare("INSERT INTO todos (id, description, completed, device_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer smtm.Close()

	_, err = smtm.Exec(todo.ID, todo.Description, todo.Completed, todo.DeviceId)
	if err != nil {
		return err
	}
	return nil
}
