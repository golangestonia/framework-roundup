package controllers

import beego "github.com/beego/beego/v2/server/web"

type PetsController struct {
	beego.Controller
}

func (c *PetsController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *PetsController) ListPets() {
	c.ServeJSON()
}

func (c *PetsController) CreatePet() {

}

func (c *PetsController) ListByStatus() {

}

func (c *PetsController) GetPet() {

}

func (c *PetsController) UpdatePet() {

}

func (c *PetsController) DeletePet() {

}
