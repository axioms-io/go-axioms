// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"go-axioms/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/permission", &controllers.PermissionController{})
	beego.Router("/private", &controllers.PrivateController{})
	beego.Router("/public", &controllers.PublicController{})
	beego.Router("/role", &controllers.RoleController{})
}
