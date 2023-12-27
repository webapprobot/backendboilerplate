package systemRoutes

import (
	physiscalBooksController "github.com/webappbot/backendboilerplate/src/controllers/physicalBook"
	"github.com/gofiber/fiber/v2"
)

func (f SystemRoutes) BookHolderLibraries(api *fiber.Group) {
	api.Post("/bookHolderLibraries", physiscalBooksController.CreatePhysicalBookHolderLibrary)
	// api.Get("/bookHolderLibraries", physiscalBooksController.ReadPhysicalBookHolderLibraries)
	api.Get("/bookHolderLibrary", physiscalBooksController.ReadPhysicalBookHolderLibrary)
	api.Get("/bookHolderLibraryBooks", physiscalBooksController.ReadPhysicalBookHolderLibraryBooks)
}
