package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about public
type PublicController struct {
	beego.Controller
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (c *PublicController) Get() {
	c.Data["json"] = map[string]string{
		"message": "Hello from a public endpoint!",
	}
	c.ServeJSON()
}
