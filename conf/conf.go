package conf

import "os"

type appConfig struct {
	Domain   string
	Audience string
}

// App global variables for env variables
var App = &appConfig{
	Domain:   os.Getenv("AXIOMS_DOMAIN"),
	Audience: os.Getenv("AXIOMS_AUDIENCE"),
}
