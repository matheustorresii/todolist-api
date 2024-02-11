package handler

import (
	"database/sql"
	"todolist/internal/app/model"

	"github.com/gofiber/fiber/v2"
)

func GetTodo(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		deviceID := c.Query("device_id")

		todos, err := getTodosByDeviceID(db, deviceID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(todos)
	}
}

func getTodosByDeviceID(db *sql.DB, deviceID string) ([]model.Todo, error) {
	todos := []model.Todo{}

	rows, err := db.Query("SELECT id, description, completed, device_id FROM todos WHERE device_id = ?", deviceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.ID, &todo.Description, &todo.Completed, &todo.DeviceId); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
