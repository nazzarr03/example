package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nazzarr03/example/routes"
)

func main() {
	app := fiber.New() // Fiber uygulamasını başlatıyoruz.

	routes.Setup(app)   // routes klasöründeki Setup fonksiyonunu çağırıyoruz.
	app.Listen(":8080") // Uygulamaya 8080 portundan erişebiliriz.
}
