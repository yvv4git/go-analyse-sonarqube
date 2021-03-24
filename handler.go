package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type sumHandler struct{}

// sum is used as handler action.
func (h *sumHandler) sum(c *fiber.Ctx) error {
	x, err := strconv.Atoi(c.Params("x"))
	if err != nil {
		return err
	}

	y, err := strconv.Atoi(c.Params("y"))
	if err != nil {
		return err
	}

	sum := x + y

	return c.JSON(fiber.Map{
		"sum": sum,
	})
}

func SetupWebApp() *fiber.App {
	sumController := &sumHandler{}

	app := fiber.New()
	app.Get("sum/:x/:y", sumController.sum)

	return app
}
