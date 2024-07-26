package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nazzarr03/example/controllers"
)

func Setup(app *fiber.App) {
	// controllers dosyasında yazdığımız endpointleri burada çağırıyoruz.
	app.Get("books", controllers.GetBooks)
	app.Get("books/:id", controllers.GetBookByID)
	app.Post("books", controllers.CreateBook)
}
