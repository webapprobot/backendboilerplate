package main

import (
	"github.com/webappbot/backendboilerplate/cmd"
	// "fmt"
	// _ "github.com/webappbot/backendboilerplate/docs"
	// "github.com/webappbot/backendboilerplate/routes"
	// swagger "github.com/arsmn/fiber-swagger/v2"
	// "github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
)

// @title backendboilerplate API
// @version 1.0
// @description backendboilerplate API
// @termsOfService http://backendboilerplate.com
// @contact.name Brian Onang'o
// @contact.email brian@backendboilerplate.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/

func main() {
	cmd.Execute()
}
