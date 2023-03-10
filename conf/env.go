package conf

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SecretKey     string `env:"SW_SECRET_KEY"`
	DomainName    string `env:"DOMAIN_NAME"`
	SubdomainName string `env:"SUBDOMAIN_NAME"`
}

func Loadenv() (*Config, *error) {
	var config = &Config{}

    _, err := os.Stat(".env")

    if !os.IsNotExist(err) {
	    error := godotenv.Load(".env")
        if error != nil {
            return config, &error
        }
    }

	config.SecretKey = os.Getenv("SW_SECRET_KEY")
	config.DomainName = os.Getenv("DOMAIN_NAME")
	config.SubdomainName = os.Getenv("SUBDOMAIN_NAME")

	return config, nil
}
