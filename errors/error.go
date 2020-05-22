package errors

// AxiomsError is a custom error
func AxiomsError(text map[string]string, code string) error {
	return &errorString{text, code}
}

type errorString struct {
	name map[string]string
	code string
}

func (e *errorString) Error() string {
	return e.code
}
