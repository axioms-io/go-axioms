package errors

import (
	"net/http"

	"github.com/astaxie/beego"
)

// ErrorController namespace controller
type ErrorController struct {
	beego.Controller
}

// Error401 handler
func (e *ErrorController) Error401() {
	var description string = e.Ctx.Input.Param(":all")
	e.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	e.Data["json"] = AxiomsError(
		"unauthorized_access",
		description,
		http.StatusUnauthorized,
	)
	e.ServeJSON()
}

// Error403 handler
func (e *ErrorController) Error403() {
	e.Ctx.ResponseWriter.WriteHeader(http.StatusForbidden)
	e.Data["json"] = AxiomsError(
		"insufficient_permission",
		"Insufficient role, scope or permission",
		http.StatusForbidden,
	)
	e.ServeJSON()
}

// Error404 handler
// func (e *ErrorController) Error404() {
// 	e.Ctx.ResponseWriter.WriteHeader(404)
// 	e.Data["json"] = AxiomsError(
// 		"page_not_found",
// 		"This route does not exist",
// 		404,
// 	)
// 	e.ServeJSON()
// }

// AxiomsError is a custom error
func AxiomsError(err string, description string, code int) error {
	var errObj = map[string]string{
		"error":             err,
		"error_description": description,
	}
	return &errorResponse{errObj, code}
}

// ErrorResponse exported
type errorResponse struct {
	Name map[string]string `json:"errmsg"`
	Code int               `json:"errcode"`
}

func (e *errorResponse) Error() string {
	return e.Name["error_description"]
}
