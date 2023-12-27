package routes

import (
	"fmt"
	"net"
	"os"
	"reflect"

	"github.com/webappbot/backendboilerplate/config"
	"github.com/webappbot/backendboilerplate/docs"
	systemRoutes "github.com/webappbot/backendboilerplate/routes/systemRoutes"
	"github.com/webappbot/backendboilerplate/src/controllers"
	authControllers "github.com/webappbot/backendboilerplate/src/controllers/auth"
	booksController "github.com/webappbot/backendboilerplate/src/controllers/books"
	charitiesController "github.com/webappbot/backendboilerplate/src/controllers/charities"
	librariesController "github.com/webappbot/backendboilerplate/src/controllers/library"
	physiscalBooksController "github.com/webappbot/backendboilerplate/src/controllers/physicalBook"
	shelvesController "github.com/webappbot/backendboilerplate/src/controllers/shelf"
	validationControllers "github.com/webappbot/backendboilerplate/src/controllers/validator"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
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

	// programmatically set swagger info
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
	api.Get("/dev", authControllers.Status)

	// Authentication end points
	noauth := api.Group("/noauth")
	// auth.Post("/login", controllers.LoginUser)
	noauth.Post("/user", authControllers.CreateUser)
	noauth.Get("/activate", authControllers.ActivateUser)
	noauth.Post("/resetPassword", authControllers.CreatePassword)
	noauth.Post("/forgotPassword", authControllers.ForgotPassword)

	auth := api.Group("/auth")
	auth.Post("/login", authControllers.Login)
	auth.Get("/token", authControllers.Token)
	auth.Get("/social/github", authControllers.Githublogin)
	auth.Get("/social/linkedIn", authControllers.LinkedInlogin)
	auth.Get("/social/facebook", authControllers.Facebooklogin)
	auth.Get("/social/twitter", authControllers.Twitterlogin)

	books := api.Group("/books")
	books.Post("/subject", booksController.CreateSubject)
	books.Patch("/subject", booksController.UpdateSubject)
	books.Get("/subjects", booksController.ReadSubject)
	books.Delete("/subjects", booksController.DeleteSubject)
	books.Post("/grade", booksController.CreateGrade)
	books.Patch("/grade", booksController.UpdateGrade)
	books.Get("/grades", booksController.ReadGrade)
	books.Delete("/grades", booksController.DeleteGrade)
	books.Post("/size", booksController.CreateSize)
	books.Patch("/size", booksController.UpdateSize)
	books.Get("/sizes", booksController.ReadSize)
	books.Delete("/sizes", booksController.DeleteSize)
	books.Post("/color", booksController.CreateColor)
	books.Patch("/color", booksController.UpdateColor)
	books.Get("/colors", booksController.ReadColor)
	books.Delete("/colors", booksController.DeleteColor)
	books.Post("/cover", booksController.CreateCover)
	books.Patch("/cover", booksController.UpdateCover)
	books.Get("/covers", booksController.ReadCover)
	books.Delete("/covers", booksController.DeleteCover)
	books.Post("/publisher", booksController.CreatePublisher)
	books.Patch("/publisher", booksController.UpdatePublisher)
	books.Get("/publishers", booksController.ReadPublisher)
	books.Delete("/publishers", booksController.DeletePublisher)
	books.Post("/author", booksController.CreateAuthor)
	books.Patch("/author", booksController.UpdateAuthor)
	books.Get("/authors", booksController.ReadAuthor)
	books.Delete("/authors", booksController.DeleteAuthor)
	books.Post("/binding", booksController.CreateBinding)
	books.Patch("/binding", booksController.UpdateBinding)
	books.Get("/bindings", booksController.ReadBinding)
	books.Delete("/bindings", booksController.DeleteBinding)
	book := api.Group("/book")
	book.Post("/", booksController.CreateBook)
	book.Patch("/", booksController.UpdateBook)
	book.Get("/", booksController.GetBook)
	book.Delete("/", booksController.DeleteBook)

	api.Post("/charity", charitiesController.CreateCharity)
	api.Patch("/charity", charitiesController.UpdateCharity)
	api.Get("/charities", charitiesController.ReadCharities)
	api.Delete("/charities", charitiesController.DeleteCharity)

	// shelves
	api.Post("/shelf", shelvesController.CreateShelf)
	api.Patch("/shelf", shelvesController.UpdateShelf)
	api.Get("/shelves", shelvesController.ReadShelves)
	api.Delete("/shelves", shelvesController.DeleteShelf)

	// shelf owners
	api.Post("/shelfowner", shelvesController.CreateShelfOwner)
	api.Patch("/shelfowner", shelvesController.UpdateShelfOwner)
	api.Get("/shelfowners", shelvesController.ReadShelfOwners)
	api.Delete("/shelfowners", shelvesController.DeleteShelfOwner)

	// library owners
	api.Post("/libraryowner", librariesController.CreateLibraryOwner)
	api.Patch("/libraryowner", librariesController.UpdateLibraryOwner)
	api.Get("/libraryowners", librariesController.ReadLibraryOwners)
	api.Delete("/libraryowners", librariesController.DeleteLibraryOwner)

	// libraries
	api.Post("/library", librariesController.CreateLibrary)
	api.Patch("/library", librariesController.UpdateLibrary)
	api.Get("/libraries", librariesController.ReadLibraries)
	api.Delete("/libraries", librariesController.DeleteLibrary)

	api.Post("/bookOwnerLibraries", physiscalBooksController.CreatePhysicalBookOwnerLibrary)
	// api.Get("/bookOwnerLibraries", physiscalBooksController.ReadPhysicalBookOwnerLibraries)
	api.Get("/bookOwnerLibrary", physiscalBooksController.ReadPhysicalBookOwnerLibrary)
	api.Get("/bookOwnerLibraryBooks", physiscalBooksController.ReadPhysicalBookOwnerLibraryBooks)

	// User related end points
	user := api.Group("/user")
	user.Get("/:id", authControllers.GetUser)
	user.Patch("/:id", authControllers.UpdateUser)
	user.Delete("/:id", authControllers.DeleteUser)

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
