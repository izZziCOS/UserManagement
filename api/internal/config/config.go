// Package config provides configuration utilities for the application
package config

import "os"

// BuildDSN constructs a PostgreSQL database connection string (DSN) from environment variables
func BuildDSN() string {
	return "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable TimeZone=UTC"
}
