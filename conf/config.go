package conf

import "github.com/sakirsensoy/genv"

type appConfig struct {
	Domain   string
	Audience string
}

var App = &appConfig{
	Domain:   genv.Key("AXIOMS_DOMAIN").String(),
	Audience: genv.Key("AXIOMS_AUDIENCE").String(),
}
