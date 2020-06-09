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
	axerr "go-axioms/errors"
	"go-axioms/filters"

	"github.com/astaxie/beego"
)

// func methodSwitcher(in string) {
// 	switch method := in; method {
// 	case "GET":
// 		perm := []string{"sample:read"}
// 		return filters.HasRequiredPermissions(perm)
// 	case "POST":
// 		perm := []string{"sample:create"}
// 		return filters.HasRequiredPermissions(perm)
// 	case "PUT":
// 		perm := []string{"sample:update"}
// 		return filters.HasRequiredPermissions(perm)
// 	default:
// 		perm := []string{"sample:unknown"}
// 		return filters.HasRequiredPermissions(perm)
// 	}
// }

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/permission", &controllers.PermissionController{})
	beego.InsertFilter(
		"/permission",
		beego.BeforeExec,
		filters.HasRequiredPermissions([]string{"sample:read"}),
	)

	beego.Router("/private", &controllers.PrivateController{})
	beego.Router("/public", &controllers.PublicController{})
	beego.Router("/role", &controllers.RoleController{})

	beego.Router("/401/:all", &axerr.ErrorController{}, "GET:Error401")
	beego.Router("/403", &axerr.ErrorController{}, "GET:Error403")
}
