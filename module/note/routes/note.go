package noteRoutes

import (
	noteController "golang-boilerplate-example/module/note/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/notes")
	
	// Read all Notes
	note.Get("/", noteController.GetAllNotes)
	// Read one Note
	note.Get("/:id", noteController.GetNote)
	// Create a Note
	note.Post("/", noteController.CreateNotes)
	// Update one Note
	note.Put("/:id", noteController.UpdateNote)
	// Delete one Note
	note.Delete("/:id", noteController.DeleteNote)
}
