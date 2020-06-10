package controllers

import (
	"github.com/astaxie/beego"
)

// IndexController '/route' namespace controller
type IndexController struct {
	beego.Controller
}

// Get method for index route
func (i *IndexController) Get() {
	i.Data["json"] = map[string]string{
		"api": "Beego Sample APIs",
	}
	i.ServeJSON()
}
