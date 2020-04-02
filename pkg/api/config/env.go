package config

import (
	"os"
)

// ServiceConf holds the config values for the REST API service.
type ServiceConf struct {
	Address string
	Port    string
}

// DBConf struct holds the config for the acces of the target database.
type DBConf struct {
	Type     string
	Address  string
	Port     string
	Username string
	Password string
}

// getEnv looks up if the environment variable is present and if yes returns the value. If not it returns default value.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetServiceConf fills the ServoiceConf struct with environemnt or default values if env vars are not present.
func (sc *ServiceConf) GetServiceConf() {
	// var sc ServiceConf

	sc.Address = getEnv("API_BIND", "0.0.0.0")
	sc.Port = getEnv("API_PORT", "8080")

	// return &sc
}

// GetDBConf gets all necessary config values from environment variables and stores them in an EnvVars struct.
func (dbc *DBConf) GetDBConf() {
	// var dbc DBConf

	dbc.Type = getEnv("DB_TYPE", "")
	dbc.Address = getEnv("DB_ADDRESS", "")
	dbc.Port = getEnv("DB_PORT", "")
	dbc.Username = getEnv("DB_USER", "")
	dbc.Password = getEnv("DB_PWD", "")

	// return &dbc
}
