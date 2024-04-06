package user
import (
	"github.com/gofiber/fiber/v2"
	"encoding/json"
)

//Criação da struct de Usuário
type User struct{
	FirstName 	string 	`json: "fistname"`
	LastName 	string 	`json: "lastname"`
	Email 		string 	`json: "email"`
	Age			int		`json: "age"`
	Id			int 	`json:"id"`
}

//Criação da lista de usuários
var users []*User

//Função que irá retornar no formato JSON todas as tuplas de usuários cadastrados
func GetAllUser(c *fiber.Ctx) error{
	return c.JSON(users)
}
//Função que irá retornar no formato JSON todas a tupla do usuário com base no id do mesmo
func GetUserById(c *fiber.Ctx) error{
	id, _ := c.ParamsInt("id")
	var userFound User
	
	for _, user := range users {
        if user.Id == id {
            userFound = *user
        }
    }

	return c.JSON(userFound)
}

//Função que envia os dados de adição de um novo usuário, cadastro
func PostUser(c *fiber.Ctx) error{
	type Request struct{

	}
	var newUser User
	if err:= c.BodyParser(&newUser); err != nil{
		return c.Status(500).SendString(err.Error())
	}
	users = append(users, &newUser)
	return c.JSON(newUser)
}

//Método para serializar um usuário em JSON
func (u *User) Serialize() ([]byte, error) {
    return json.Marshal(u)
}

//Método para desserializar um usuário em JSON
func DeserializeUser(data []byte) (*User, error) {
    var user User
    err := json.Unmarshal(data, &user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}