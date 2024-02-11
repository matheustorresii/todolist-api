package handler

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func DeleteTodo(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		todoId := c.Query("id")

		if err := deleteById(db, todoId); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.SendString("Deleted")
	}
}

func deleteById(db *sql.DB, id string) error {
	smtm, err := db.Prepare("DELETE FROM todos WHERE id = ?")
	if err != nil {
		return err
	}
	defer smtm.Close()

	_, err = smtm.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
