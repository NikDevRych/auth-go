package config

import "os"

type Config struct {
	ConnectionString string
	JWTSecretKey     string
}

func NewConfig() *Config {
	return &Config{
		ConnectionString: os.Getenv("connection_string"),
		JWTSecretKey:     os.Getenv("jwt_key"),
	}
}
