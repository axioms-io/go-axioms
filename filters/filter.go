package filters

import (
	"fmt"
	"go-axioms/conf"
	"go-axioms/tokens"
	"log"
	"os"

	"go-axioms/errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gopkg.in/oleiade/reflections.v1"
)

// HasRequiredScopes checks scopes
func HasRequiredScopes(viewRoles []string) beego.FilterFunc {
	return func(ctx *context.Context) {
		payload := ctx.Request.auth_jwt
		if payload == nil {
			return "", err.AxiomsError(
				"unauthorized_access",
				"Invalid Authorisation Token",
				401,
			)
		}
		if tokens.CheckScopes(payload.scope, required_scopes[0]) {
			return
		}
		return "", err.AxiomsError(
			"insufficient_permission",
			"Insufficient role, scope or permission",
			403,
		)
	}
}

// HasRequiredRoles checks roles
func HasRequiredRoles(viewRoles []string) beego.FilterFunc {
	return func(ctx *context.Context) {
		payload := ctx.Request.auth_jwt
		if payload == nil {
			return "", err.AxiomsError(
				"unauthorized_access",
				"Invalid Authorisation Token",
				401,
			)
		}
		if tokens.CheckRoles(tokenRoles, viewRoles[0]) {
			return
		}
		return "", err.AxiomsError(
			"insufficient_permission",
			"Insufficient role, scope or permission",
			403,
		)
	}
}

// HasRequiredPermissions checks permissions
func HasRequiredPermissions(viewPermissions []string) beego.FilterFunc {
	return func(ctx *context.Context) {
		payload, err := reflections.GetField(ctx.Request, "auth_jwt")
		if err != nil {
			return "", err.AxiomsError(
				"unauthorized_access",
				"Invalid Authorisation Token",
				401,
			)
		}
		token_permissions, err = reflections.GetField(
			payload,
			fmt.Sprintf("https://%s/claims/permissions", os.Getenv("AXIOMS_DOMAIN")))
		if tokens.CheckPermissions(tokenPermissions, viewPermissions[0]) {
			return
		}
		return "", err.AxiomsError(
			"insufficient_permission",
			"Insufficient role, scope or permission",
			403)
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
		return "", errors.AxiomsError(
			"unauthorized_access",
			"Invalid Authorization Token",
			401,
		)
	}
}
