package tokens

import (
	"strings"

	"github.com/fatih/set"
	jose "gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func hasBearerToken(reqObj jwt.JSONWebToken) error {
	var headerName string = "Authorization"
	var tokenPrefix string = "bearer"
	var err error
	var authHeader jose.Header
	for i, s := range reqObj.Headers {
		if s.KeyID == headerName {
			authHeader = s
		}
	}
	if authHeader.KeyID != headerName {
		var errObj = map[string]string{
			"error":             "unauthorised_access",
			"error_description": "Missing Authorisation Header",
		}
		return err.AxiomsError(errObj, 401)
	}
	return nil
}

func hasValidToken(token jwt.JSONWebToken) {

}

func checkTokenValidity(token jwt.JSONWebToken, key string) {

}

func getPayloadFromToken(jwt.JSONWebToken) {

}

func checkScopes(providedScopes string, requiredScopes []string) bool {
	if len(requiredScopes) == 0 {
		return true
	}
	var tmp []string = strings.Split(providedScopes, " ")
	var tokenScopes = set.New(set.ThreadSafe)
	for i, s := range tmp {
		tokenScopes.Add(s)
	}
	scopes := set.New(set.ThreadSafe)
	for i, s := range requiredScopes {
		scopes.Add(s)
	}
	return set.Intersection(tokenScopes, scopes).Size() > 0
}

func checkRoles(tokenRoles []string, viewRoles []string) bool {
	if len(viewRoles) == 0 {
		return true
	}
	token := set.New(set.ThreadSafe)
	for i, s := range tokenRoles {
		token.Add(s)
	}
	views := set.New(set.ThreadSafe)
	for i, s := range viewRoles {
		views.Add(s)
	}
	return set.Intersection(token, views).Size() > 0
}

func checkPermissions() {

}

func getKeyFromJWKSjson() {

}
