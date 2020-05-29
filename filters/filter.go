package filters

import (
	token "go-axioms/tokens"
)

func hasRequiredScopes(viewRoles []string) {
	payload := getattr(request, "auth_jwt", None)
	if payload == nil {
		return "", err.AxiomsError(
			"unauthorized_access",
			"Invalid Authorisation Token", 
			401,
		)
	}
	if token.checkScopes(payload.scope, required_scopes[0]) {
		return fn(*args, **kwargs)
	}
	return "", err.AxiomsError(
		"insufficient_permission",
		"Insufficient role, scope or permission", 
		403,
	)
}

func hasRequiredRoles(viewRoles []string) {
	payload := getatrr(request, "auth_jwt", nil)
	if payload == nil {
		return "", err.AxiomsError(
			"unauthorized_access",
			"Invalid Authorisation Token", 
			401,
		)
	}
	if token.checkRoles(tokenRoles, viewRoles[0]) {
		return fn(*args, **kwargs)
	}
	return "", err.AxiomsError(
		"insufficient_permission",
		"Insufficient role, scope or permission", 
		403,
	)
}

func hasRequiredPermissions(viewPermissions []string) {
	payload := getatrr(request, "auth_jwt", nil)
	if payload == nil {
		return "", err.AxiomsError(
			"unauthorized_access",
			"Invalid Authorisation Token", 
			401,
		)
	}
	var token_permissions []string
	token_permissions = getattr(
		payload,
		"https://{}/claims/permissions".format(app.config["AXIOMS_DOMAIN"]),
		[]
	)
	if token.checkPermissions(tokenPermissions, viewPermissions[0]) {
		return fn(*args, **kwargs)
	}
	return "", err.AxiomsError(
		"insufficient_permission",
		"Insufficient role, scope or permission", 
		403)
}

func has_valid_access_token() {
	try:
		app.config["AXIOMS_DOMAIN"]
		app.config["AXIOMS_AUDIENCE"]
	except KeyError as e:
		raise Exception(
			"🔥🔥 Please set value for {} in a .env file. For more details review axioms-flask-py docs.".format(
				e
			)
		)
	token = has_bearer_token(request)
	if token and has_valid_token(token):
		return fn(*args, **kwargs)
	else:
		raise AxiomsError(
			{
				"error": "unauthorized_access",
				"error_description": "Invalid Authorization Token",
			},
			401,
		)
}
