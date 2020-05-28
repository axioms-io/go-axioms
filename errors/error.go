package errors

// AxiomsError is a custom error
func AxiomsError(err string, description string, code int) error {
	var errObj = map[string]string{
		"error":             err,
		"error_description": description,
	}
	return &errorString{errObj, code}
}

type errorString struct {
	name map[string]string
	code int
}

func (e *errorString) Error() string {
	return e.name["error"]
}
