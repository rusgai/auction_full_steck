package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.JSON(fiber.Map{
			"status":  200,
			"message": "Home api massage",
		})
		if err != nil {
			return err
		}
		return nil
	})

	app.Listen(":8001")
}
