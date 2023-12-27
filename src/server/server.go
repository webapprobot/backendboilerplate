package server

import (
	"fmt"
	"log"

	"github.com/webappbot/backendboilerplate/config"
	_ "github.com/webappbot/backendboilerplate/docs"
	"github.com/webappbot/backendboilerplate/routes"
	"github.com/webappbot/backendboilerplate/src/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Start(production bool) {
	startMQTT()
	env := "dev"
	if production {
		env = "prod"
	}

	port, err := config.GetConfig("PORT", env)
	if err != nil {
		log.Fatal(err)
	}
	portStr := fmt.Sprintf("%v", port)
	portStr = fmt.Sprintf(":%s", portStr)
	db.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)
	if err = app.Listen(portStr); err != nil {
		log.Fatal(err)
	}
}

func startMQTT() {
	// We don't need this now
	// Exec(func() interface{} {
	// 	mqttConnect()
	// 	return nil
	// })
}
