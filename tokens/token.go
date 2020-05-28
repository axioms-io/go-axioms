package tokens

import (
	// Standard Imports
	"strings"
	"time"
	"net/http"
	"encoding/json"
	"io/ioutil"
	
	// SDK Imports
	axerr "go-axioms/errors"

	// Package Imports
	"github.com/fatih/set"
	jose "gopkg.in/square/go-jose.v2"
	"github.com/eko/gocache/store"
)

memcacheStore := store.NewMemcache(
	memcache.New("localhost"),
	&store.Options{
		Expiration: 300 * time.Second
	}
)
cacheManager := cache.New(memcacheStore)

func hasBearerToken(headers map[string]interface{}) (string, error) {
	var headerName string = "Authorization"
	var tokenPrefix string = "bearer"
	var authHeader interface{}
	if headers[headerName] != nil {
		authHeader = headers[headerName]
	} else {
		return "", axerr.AxiomsError(
			"unauthorized_access",
			"Missing Authorisation Header", 
			"401"
		)
	}
	// NOTE: What is part of the interface that makes the value of the header?
	split := strings.Split(authHeader, " ")
	bearer, token := split[0], split[1]

	if strings.ToLower(bearer) == tokenPrefix && token != "" {
		return token, nil
	} else {
		var errObj = map[string]string{
			return "", axerr.AxiomsError(
				"unauthorized_access",
				"Invalid Authorisation Bearer", 
				"401"
			)
	}

	return "", nil
}

func hasValidToken(token jwt.JSONWebToken) {

}

func checkTokenValidity(token string, key interface{}) string {
	payload, err := getPayloadFromToken(token, key)
	now := time.Now().Unix()
	if payload == "" && now <= payload.exp {
		return payload
	}
	return ""
}

func getPayloadFromToken(token string, key interface{}) (string, error) {
	payload, headers, err := jose.Decode(token, key)
	if err != nil {
		return "", err
	}
	return payload, nil
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
	view := set.New(set.ThreadSafe)
	for i, s := range viewRoles {
		view.Add(s)
	}
	return set.Intersection(token, view).Size() > 0
}

func checkPermissions(tokenPermissions []string, viewPermissions []string) bool {
	if len(viewPermissions) == 0 {
		return true
	}
	token := set.New(set.ThreadSafe)
	for i, s := range tokenPermissions {
		token.Add(s)
	}
	view := set.New(set.ThreadSafe)
	for i, s := range viewPermissions {
		view.Add(s)
	}
	return set.Intersection(token, view).Size() > 0
}

func getKeyFromJWKSjson(tenant string, kid string) (jose.JSONWebKeySet, err) {
	data = cacheFetch("https://" + tenant + "/oauth2/.well-known/jwks.json", 600)
	key = &jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{},
	}
	err := json.Unmarshal([]byte(data), key)
	if err != nil {
		return key, axerr.AxiomsError(
			"unathorized_access",
			"Invalid Access Token",
			"401"
		)
	}
	return key, nil
}

func cacheFetch(url string, timeOfLife int) {
	cached, err = cacheManager.Get("jwks" + url)
	if err != nil {
		return cached
	}
	response, err := http.Get(url)
	data, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	err := cacheManager.Set("jwks" + url, data, &cache.Options{Expiration: timeOfLife*time.Second})
	return data
}
