package controllers

import "github.com/gofiber/fiber/v2"

// Status godoc
// @Summary Returns a 200
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /dev [get]
// @Tags   dev
func Status(c *fiber.Ctx) error {
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
