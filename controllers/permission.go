package controllers

import (
	"github.com/astaxie/beego"
)

// Operations for /permission
type PermissionController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (c *PermissionController) Post() {
	c.Data["json"] = map[string]string{
		"message": "Sample created.",
	}
	c.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (c *PermissionController) Patch() {
	c.Data["json"] = map[string]string{
		"message": "Sample updated.",
	}
	c.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (c *PermissionController) Get() {
	c.Data["json"] = map[string]string{
		"message": "Sample read.",
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (c *PermissionController) Delete() {
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
// func (u *PermissionController) Login() {
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
// func (u *PermissionController) Logout() {
// 	u.Data["json"] = "logout success"
// 	u.ServeJSON()
// }
