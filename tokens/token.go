package tokens

import (
	"strings"
	"time"

	"github.com/fatih/set"
	jose "gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func hasBearerToken(reqObj jwt.JSONWebToken) (string, error) {
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
		return "", err.AxiomsError(errObj, 401)
	}
	split := strings.Split(authHeader.KeyID, " ")
	bearer, token := split[0], split[1]

	if strings.ToLower(bearer) == tokenPrefix && token != "" {
		return token, nil
	} else {
		var errObj = map[string]string{
			"error":             "unauthorised_access",
			"error_description": "Invalid Authorisation Bearer",
		}
		return "", err.AxiomsError(errObj, 401)
	}

	return "", nil
}

func hasValidToken(token jwt.JSONWebToken) {

}

func checkTokenValidity(token jwt.JSONWebToken, key string) {
	payload := getPayloadFromToken(token, key)
	now := time.Now().Unix()
	if payload != nil && now <= payload.exp {
		return payload
	}
	return nil
}

func getPayloadFromToken(token jwt.JSONWebToken, key string) {
	return 1
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

func checkPermissions(tokenPermissions []string, viewPermissions []string) bool {
	if len(viewPermissions) == 0 {
		return true
	}
	token := set.New(set.ThreadSafe)
	for i, s := range tokenPermissions {
		token.Add(s)
	}
	views := set.New(set.ThreadSafe)
	for i, s := range viewPermissions {
		views.Add(s)
	}
	return set.Intersection(token, views).Size() > 0
}

func getKeyFromJWKSjson() {

}
