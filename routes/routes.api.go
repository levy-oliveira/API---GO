package routes
import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-API/Model"
)

//Rotas que direcionam a API para realizar a requisição, dependendo do tipo de requisição demandada.
func Routers(app *fiber.App) {

	//quando entra no localhot8080, já faz a captura de todos os usuários cadastrados. Executa a função GetAllUser.
	app.Get("/", user.GetAllUser)

	//quando entra no localhot8080/:id, faz a captura do usuário com Id passado como parâmetro. Executa a função GetUserById.
	app.Get("/:id", user.GetUserById)
	
	//dada uma requisição POST, irá redireceionar à função PostUser. 
	app.Post("/", user.PostUser)
}

