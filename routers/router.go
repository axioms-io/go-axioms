// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"axioms/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// ns := beego.NewNamespace("/v1",
	// 	beego.NSNamespace("/object",
	// 		beego.NSInclude(
	// 			&controllers.ObjectController{},
	// 		),
	// 	),
	// 	beego.NSNamespace("/user",
	// 		beego.NSInclude(
	// 			&controllers.UserController{},
	// 		),
	// 	),
	// )
	perm := beego.NewNamespace("/permission",
		beego.NSInclude(
			&controllers.PermissionController{},
		),
	)
	priv := beego.NewNamespace("/private",
		beego.NSInclude(
			&controllers.PrivateController{},
		),
	)
	publ := beego.NewNamespace("/public",
		beego.NSInclude(
			&controllers.PublicController{},
		),
	)
	role := beego.NewNamespace("/role",
		beego.NSInclude(
			&controllers.RoleController{},
		),
	)
	beego.AddNamespace(perm, priv, publ, role)
}
