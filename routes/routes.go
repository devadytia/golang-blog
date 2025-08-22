package routes

import (
	"blog/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/article/:limit/:offset", handlers.GetAllPosts)
	app.Post("/article", handlers.CreatePosts)
	app.Get("/article/:id", handlers.GetPostById)
	app.Post("/article/:id", handlers.UpdatePost)
	app.Delete("/article/:id", handlers.DeletePost)
}
