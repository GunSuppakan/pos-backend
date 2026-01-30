package utility

import "github.com/gofiber/fiber/v2"

func ResponseSuccess(c *fiber.Ctx, data interface{}) error {
	if data != nil {
		return c.JSON(fiber.Map{
			"status":  200,
			"data":    data,
			"message": "success",
		})
	}
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
	})
}

func ResponseError(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  status,
		"message": message,
	})
}
