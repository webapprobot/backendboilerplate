package controllers

import (
	"fmt"

	"github.com/webappbot/backendboilerplate/config"
	"github.com/gofiber/fiber/v2"
)

var name string

func Welcome(c *fiber.Ctx) error {
	name, _ := config.GetConfig("service", "dev", "name")
	return c.JSON(fiber.Map{"status": "success", "message": fmt.Sprintf("Welcome to %s API", name), "data": nil})
}
