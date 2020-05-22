package controllers

import (
	"github.com/astaxie/beego"
)

// PrivateController ...
// Operations for /private endpoints
type PrivateController struct {
	beego.Controller
}

// Get ...
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (c *PrivateController) Get() {
	c.Data["json"] = map[string]string{
		"message": "All good. You are authenticated!",
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
// func (u *PrivateController) Login() {
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
// func (u *PrivateController) Logout() {
// 	u.Data["json"] = "logout success"
// 	u.ServeJSON()
// }
