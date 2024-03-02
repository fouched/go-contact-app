package main

import "github.com/alexedwards/scs/v2"

type AppConfig struct {
	InProduction bool
	Session      *scs.SessionManager
}
