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

var perms = map[string][]string{
	"GET":    {"sample:read"},
	"PUT":    {"sample:update"},
	"POST":   {"sample:create"},
	"DELETE": {"sample:remove"},
}

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/permission", &controllers.PermissionController{})
	beego.InsertFilter("/permission", beego.BeforeExec, filters.HasValidAccessToken())
	beego.InsertFilter("/permission", beego.BeforeExec, filters.HasRequiredPermissions(perms))

	beego.Router("/private", &controllers.PrivateController{})
	beego.InsertFilter("/permission", beego.BeforeExec, filters.HasValidAccessToken())
	beego.InsertFilter("/permission", beego.BeforeExec, filters.HasRequiredScopes([]string{"openid", "profile"}))

	beego.Router("/public", &controllers.PublicController{})

	beego.Router("/role", &controllers.RoleController{})
	beego.InsertFilter("/permission", beego.BeforeExec, filters.HasValidAccessToken())
	beego.InsertFilter("/permission", beego.BeforeExec, filters.HasRequiredRoles([]string{"sample:role"}))

	beego.Router("/401/:all", &axerr.ErrorController{}, "GET:Error401")
	beego.Router("/403", &axerr.ErrorController{}, "GET:Error403")
}
