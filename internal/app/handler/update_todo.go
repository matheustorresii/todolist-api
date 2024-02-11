package handler

import (
	"database/sql"
	"todolist/internal/app/model"

	"github.com/gofiber/fiber/v2"
)

func UpdateTodo(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		todoId := c.Query("id")

		var todo model.Todo
		if err := c.BodyParser(&todo); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		if err := updateTodoById(db, todoId, &todo); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.SendString("Updated")
	}
}

func updateTodoById(db *sql.DB, id string, todo *model.Todo) error {
	smtm, err := db.Prepare("UPDATE todos SET description = ?, completed = ?, device_id = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer smtm.Close()

	_, err = smtm.Exec(todo.Description, todo.Completed, todo.DeviceId, id)
	if err != nil {
		return err
	}

	return nil
}
