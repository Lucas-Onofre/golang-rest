package handlers

import (
	"github.com/Lucas-Onofre/golang-rest/database"
	"github.com/Lucas-Onofre/golang-rest/models"
	"github.com/gofiber/fiber/v2"
)

func ListTasks(c *fiber.Ctx) error {
	tasks := []models.Task{}
	database.DB.Db.Find(&tasks)

	return c.Status(200).JSON(tasks)
}

func GetTaskById(c *fiber.Ctx) error {
	id := c.Params("id")
	task := new(models.Task)

	database.DB.Db.Find(&task, id)

	return c.Status(200).JSON(task)
}

func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error(),
        })
	}

	database.DB.Db.Create(&task)

	return c.Status(200).JSON(task)
}

func DeleteTaskById(c *fiber.Ctx) error {
	id := c.Params("id")
	task := new(models.Task)

	database.DB.Db.Find(&task, id)
	if task.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Task not found",
		})
	}

	database.DB.Db.Delete(&task)

	return c.Status(200).JSON(fiber.Map{
		"message": "Task deleted",
	})
}
