package filters

import (
	"fmt"
	"go-axioms/conf"
	"go-axioms/tokens"
	"log"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gopkg.in/oleiade/reflections.v1"
)

// HasRequiredScopes checks scopes
func HasRequiredScopes(viewScopes []string) beego.FilterFunc {
	return func(ctx *context.Context) {
		payload, err := reflections.GetField(ctx.Request, "auth_jwt")
		if err != nil {
			ctx.Redirect(http.StatusUnauthorized, "/401/Invalid_Authorization_Token'")
		}
		if tokens.CheckScopes(payload.scope, viewScopes) {
			return
		}
		ctx.Redirect(http.StatusUnauthorized, "/403")
	}
}

// HasRequiredRoles checks roles
func HasRequiredRoles(viewRoles []string) beego.FilterFunc {
	return func(ctx *context.Context) {
		payload, err := reflections.GetField(ctx.Request, "auth_jwt")
		if err != nil {
			ctx.Redirect(http.StatusUnauthorized, "/401/Invalid_Authorization_Token'")
		}
		tokenRoles, err := reflections.GetField(
			payload,
			fmt.Sprintf("https://%s/claims/roles", conf.App.Domain),
		)
		if tokens.CheckRoles(tokenRoles, viewRoles) {
			return
		}
		ctx.Redirect(http.StatusForbidden, "/403")
	}
}

// HasRequiredPermissions checks permissions
func HasRequiredPermissions(viewPermissions map[string][]string) beego.FilterFunc {
	return func(ctx *context.Context) {
		payload, err := reflections.GetField(ctx.Request, "auth_jwt")
		if err != nil {
			ctx.Redirect(http.StatusUnauthorized, "/401/Invalid_Authorization_Token'")
		}
		val, err := reflections.GetField(
			payload,
			fmt.Sprintf("https://%s/claims/permissions", conf.App.Domain),
		)
		tokenPermissions := fmt.Sprintf("%v", val)
		if tokens.CheckPermissions(tokenPermissions, viewPermissions[ctx.Request.Method]) {
			return
		}
		ctx.Redirect(http.StatusUnauthorized, "/403")
	}
}

// HasValidAccessToken checks valid access token
func HasValidAccessToken() beego.FilterFunc {
	return func(ctx *context.Context) {
		if conf.App.Audience == "" || conf.App.Domain == "" {
			log.Panicf("🔥🔥 Please set value for {} in a .env file. For more details review axioms-flask-py docs.")
		}
		token, err := tokens.HasBearerToken(ctx.Request)
		if err != nil && tokens.HasValidToken(token) {
			return
		}
		ctx.Redirect(http.StatusUnauthorized, "/401/"+err.Error())
	}
}
