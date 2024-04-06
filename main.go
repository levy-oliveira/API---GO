package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-API/routes"
)

func main() {
	//Criação de uma nova instância no Fiber
	app := fiber.New()

	//Designação das rotas que serão executadas pela API(pasta routes)
	routes.Routers(app)

	//Inicia o servido na porta localhost:8080
	app.Listen(":8080")
}

