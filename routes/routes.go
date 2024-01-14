package routes

import (
	"fmt"
	"net"
	"os"
	"reflect"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/webappbot/backendboilerplate/config"
	"github.com/webappbot/backendboilerplate/docs"

	systemRoutes "github.com/webappbot/backendboilerplate/local-src/systemRoutes"
	controllers "github.com/webappbot/backendboilerplate/src/controllers"

	systemControllers "github.com/webappbot/backendboilerplate/src/controllers"
	validationControllers "github.com/webappbot/backendboilerplate/src/helpers/validator"
)

func Setup(app *fiber.App) {
	swaggerTitle, _ := config.GetConfig("service", "dev", "displayname")
	swaggerDescription, _ := config.GetConfig("service", "dev", "description")
	domain, _ := config.GetConfig("domain")
	env := os.Getenv("env")
	port, _ := config.GetConfig("port", env)
	domainStr := fmt.Sprintf("%v", domain)
	portStr := fmt.Sprintf("%v", port)
	if net.ParseIP(domainStr) != nil || domainStr == "localhost" { // is IP address rather than domain
		domainStr += ":" + portStr
	}

	// set swagger info
	docs.SwaggerInfo.Title = swaggerTitle.(string) + " API"
	docs.SwaggerInfo.Description = swaggerDescription.(string)
	docs.SwaggerInfo.Version = config.Version()
	docs.SwaggerInfo.Host = domainStr
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	app.Get("/swagger", swagger.HandlerDefault)         // default
	app.Get("/swagger-ui", swagger.HandlerDefault)      // default
	app.Get("/swagger-ui.html", swagger.HandlerDefault) // default
	app.Get("/swagger/*", swagger.HandlerDefault)       // default

	tmpConfig := swagger.Config{ // custom
		// Title:       "swagger-ui",
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}
	app.Get("/swagger/*", swagger.New(tmpConfig))
	app.Get("/swagger-ui/*", swagger.New(tmpConfig))

	/**
	* Group is used for Routes with common prefix to define a new sub-router with optional middleware.
	**/

	// Base Api end point
	api := app.Group("/api", logger.New())
	api.Use(validationControllers.Validator)
	api.Use(recover.New())

	api.Get("/", controllers.Welcome)
	api.Get("/dev", systemControllers.Status)

	pkg := reflect.ValueOf(systemRoutes.SystemRoutes{})
	for i := 0; i < pkg.NumMethod(); i++ {
		method := pkg.Method(i)
		args := []reflect.Value{
			reflect.ValueOf(api),
		}
		_ = method.Call(args)
	}
}

// CreateApin_drugs
