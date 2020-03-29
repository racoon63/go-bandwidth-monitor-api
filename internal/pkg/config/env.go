package config

import (
	"os"
)

// EnvVars struct holds all environment values for api configuration.
type EnvVars struct {
	Database string
	Address  string
	Port     string
	Username string
	Password string
}

// Read gets all necessary config values from environment variables and stores them in an EnvVars struct.
func Read() *EnvVars {
	var conf EnvVars

	conf.Database = os.Getenv("DATABASE")
	conf.Address = os.Getenv("ADDRESS")
	conf.Port = os.Getenv("PORT")
	conf.Username = os.Getenv("USERNAME")
	conf.Password = os.Getenv("PASSWORD")

	return &conf
}
