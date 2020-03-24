package storage

import (
	"os"
)

// Envars struct holds all environment values for configuration
type Envars struct {
	Backend  string
	Username string
	Password string
}

func readEnvars() Envars {
	var e Envars
	e.Backend = os.Getenv("BACKEND")
	e.Username = os.Getenv("USERNAME")
	e.Password = os.Getenv("PASSWORD")
	return e
}
