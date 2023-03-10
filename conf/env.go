package conf

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    AccessKey string
    SecretKey string
    DomainName string
    SubdomainName string
}

func Loadenv() (*Config, *error) {
    var config = &Config{}

    error := godotenv.Load(".env")
    if error != nil {
        return config, &error
    }

    config.AccessKey = os.Getenv("SW_ACCESS_KEY")
    config.SecretKey = os.Getenv("SW_SECRET_KEY")
    config.DomainName = os.Getenv("DOMAIN_NAME")
    config.SubdomainName = os.Getenv("SUBDOMAIN_NAME")

    return config, nil
}
