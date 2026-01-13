package config

import (
	"os"
)


type Config struct {
	Name string
}


func LoadConfig() *Config {
	return &Config{
		Name: os.Getenv("NAME"),
	}
}