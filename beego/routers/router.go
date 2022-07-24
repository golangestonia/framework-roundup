package routers

import (
	"beego/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/pets", &controllers.PetsController{}, "get:ListPets")
	beego.Router("/pet", &controllers.PetsController{}, "put:CreatePet")
	beego.Router("/pet/findByStatus", &controllers.PetsController{}, "put:ListByStatus")
	beego.Router("/pet/:id:int", &controllers.PetsController{}, "get:GetPet;post:UpdatePet;delete:DeletePet")
}
