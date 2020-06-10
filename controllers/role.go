package controllers

import (
	"github.com/astaxie/beego"
)

// RoleController '/role/' namespace controller
type RoleController struct {
	beego.Controller
}

// Get method handler for /role route
func (c *RoleController) Get() {
	c.Data["json"] = map[string]string{
		"message": "Sample read.",
	}
	c.ServeJSON()
}

// Post method handler for /role route
func (c *RoleController) Post() {
	c.Data["json"] = map[string]string{
		"message": "Sample created.",
	}
	c.ServeJSON()
}

// Patch method handler for /role route
func (c *RoleController) Patch() {
	c.Data["json"] = map[string]string{
		"message": "Sample updated.",
	}
	c.ServeJSON()
}

// Delete method handler for /role route
func (c *RoleController) Delete() {
	c.Data["json"] = map[string]string{
		"message": "Sample deleted.",
	}
	c.ServeJSON()
}
