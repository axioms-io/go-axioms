package controllers

import (
	"github.com/astaxie/beego"
)

// RoleController '/route' namespace controller
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

// // @Title Login
// // @Description Logs user into the system
// // @Param	username		query 	string	true		"The username for login"
// // @Param	password		query 	string	true		"The password for login"
// // @Success 200 {string} login success
// // @Failure 403 user not exist
// // @router /login [get]
// func (u *RoleController) Login() {
// 	username := u.GetString("username")
// 	password := u.GetString("password")
// 	if models.Login(username, password) {
// 		u.Data["json"] = "login success"
// 	} else {
// 		u.Data["json"] = "user not exist"
// 	}
// 	u.ServeJSON()
// }

// // @Title logout
// // @Description Logs out current logged in user session
// // @Success 200 {string} logout success
// // @router /logout [get]
// func (u *RoleController) Logout() {
// 	u.Data["json"] = "logout success"
// 	u.ServeJSON()
// }
