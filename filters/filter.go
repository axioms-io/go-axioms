package filters

func has_required_scopes(viewRoles []string) {
	payload := getattr(request, "auth_jwt", None)
	if payload == nil {
		errObj := map[string]string{
			"error":             "unauthorized_access",
			"error_description": "Invalid Authorisation Token",
		}
		return "", err.AxiomsError(errObj, 401)
	}
	if checkScopes(payload.scope, required_scopes[0]) {
		return fn(*args, **kwargs)
	}
	errObj := map[string]string{
		"error":             "insufficient_permission",
		"error_description": "Insufficient role, scope or permission",
	}
	return "", err.AxiomsError(errObj, 403)
}

func has_required_roles(viewRoles []string) {
	payload := getatrr(request, "auth_jwt", nil)
	if payload == nil {
		errObj := map[string]string{
			"error":             "unauthorized_access",
			"error_description": "Invalid Authorisation Token",
		}
		return "", err.AxiomsError(errObj, 401)
	}
	if checkRoles(tokenRoles, viewRoles[0]) {
		return fn(*args, **kwargs)
	}
	errObj := map[string]string{
		"error":             "insufficient_permission",
		"error_description": "Insufficient role, scope or permission",
	}
	return "", err.AxiomsError(errObj, 403)
}

func has_required_permissions(viewPermissions []string) {
	payload := getatrr(request, "auth_jwt", nil)
	if payload == nil {
		errObj := map[string]string{
			"error":             "unauthorized_access",
			"error_description": "Invalid Authorisation Token",
		}
		return "", err.AxiomsError(errObj, 401)
	}
	var token_permissions []string
	token_permissions = getattr(
		payload,
		"https://{}/claims/permissions".format(app.config["AXIOMS_DOMAIN"]),
		[]
	)
	if check_permissions(tokenPermissions, viewPermissions[0]):
		return fn(*args, **kwargs)
	errObj := map[string]string{
		"error":             "insufficient_permission",
		"error_description": "Insufficient role, scope or permission",
	}
	return "", err.AxiomsError(errObj, 403)
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
