package errs

import "github.com/gofiber/fiber/v2"

func HandleHTTPError(c *fiber.Ctx, err error) error {
	if appErr, ok := err.(*AppError); ok {
		return c.Status(appErr.HTTPStatus).JSON(fiber.Map{
			"status":  appErr.Status,
			"message": appErr.Message,
		})
	}

	// fallback
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status":  ErrInternal.Status,
		"message": ErrInternal.Message,
	})
}
