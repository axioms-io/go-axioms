package tokens

import (
	// Standard Imports
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	// SDK Imports
	"go-axioms/conf"
	axerr "go-axioms/errors"

	// Package Imports
	"github.com/bluele/gcache"
	"github.com/fatih/set"
	jose "gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

var gc gcache.Cache = gcache.New(10).Build()

func HasBearerToken(reqObj *http.Request) (string, error) {
	var headerName string = "Authorization"
	var tokenPrefix string = "bearer"
	authHeader := reqObj.Header.Get(headerName)
	if authHeader == "" {
		return "", axerr.AxiomsError(
			"unauthorized_access",
			"Missing_Authorisation_Header",
			401,
		)
	}
	// What is the structure of the header?
	split := strings.Split(authHeader, " ")
	bearer, token := split[0], split[1]

	if strings.ToLower(bearer) == tokenPrefix && token != "" {
		return token, nil
	}
	return "", axerr.AxiomsError(
		"unauthorized_access",
		"Invalid_Authorisation_Bearer",
		401,
	)
}

func HasValidToken(tok string) bool {
	claims := make(map[string]interface{})
	token, _ := jwt.ParseSigned(tok)
	token.UnsafeClaimsWithoutVerification(&claims)
	key, _ := getKeyFromJWKSjson(conf.App.Domain, fmt.Sprintf("%v", claims["kid"]))
	payload, valid := checkTokenValidity(tok, key)
	if valid && payload.Audience.Contains(conf.App.Domain) {
		return true
	}
	return false
}

func checkTokenValidity(tok string, key jose.JSONWebKey) (jwt.Claims, bool) {
	payload, _ := getPayloadFromToken(tok, key)
	now := time.Now().Unix()
	if now <= payload.Expiry.Time().Unix() {
		return payload, true
	}
	return payload, false
}

func getPayloadFromToken(tok string, key jose.JSONWebKey) (jwt.Claims, error) {
	token, err := jwt.ParseSigned(tok)
	if err != nil {
		return jwt.Claims{}, err
	}
	payload := jwt.Claims{}
	if err := token.Claims(key, &payload); err != nil {
		panic(err)
	}
	return payload, nil
}

func CheckScopes(providedScopes string, requiredScopes []string) bool {
	if len(requiredScopes) == 0 {
		return true
	}
	var tmp []string = strings.Split(providedScopes, " ")
	var tokenScopes = set.New(set.ThreadSafe)
	for _, s := range tmp {
		tokenScopes.Add(s)
	}
	scopes := set.New(set.ThreadSafe)
	for _, s := range requiredScopes {
		scopes.Add(s)
	}
	return set.Intersection(tokenScopes, scopes).Size() > 0
}

func CheckRoles(tokenRoles []string, viewRoles []string) bool {
	if len(viewRoles) == 0 {
		return true
	}
	token := set.New(set.ThreadSafe)
	for _, s := range tokenRoles {
		token.Add(s)
	}
	view := set.New(set.ThreadSafe)
	for _, s := range viewRoles {
		view.Add(s)
	}
	return set.Intersection(token, view).Size() > 0
}

func CheckPermissions(tokenPermissions []string, viewPermissions []string) bool {
	if len(viewPermissions) == 0 {
		return true
	}
	token := set.New(set.ThreadSafe)
	for _, s := range tokenPermissions {
		token.Add(s)
	}
	view := set.New(set.ThreadSafe)
	for _, s := range viewPermissions {
		view.Add(s)
	}
	return set.Intersection(token, view).Size() > 0
}

func getKeyFromJWKSjson(tenant string, kid string) (jose.JSONWebKey, error) {
	data := cacheFetch("https://"+tenant+"/oauth2/.well-known/jwks.json", 600)
	key := &jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{},
	}
	err := json.Unmarshal([]byte(fmt.Sprintf("%v", data)), key)
	if err != nil {
		return key.Key(kid)[0], axerr.AxiomsError(
			"unathorized_access",
			"Invalid_Access_Token",
			401,
		)
	}
	return key.Key(kid)[0], nil
}

func cacheFetch(url string, timeOfLife int) interface{} {
	cached, err := gc.Get("jwks" + url)
	if err != nil {
		return cached
	}
	response, err := http.Get(url)
	data, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	gc.SetWithExpire("jwks"+url, data, time.Second*300)
	return data
}
