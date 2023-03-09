package conf

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    accessKey string
    secretKey string
}

func Loadenv() (*Config, *error) {
    var config = Config{}

    error := godotenv.Load(".env")
    if error != nil {
        return &config, &error
    }

    config.accessKey = os.Getenv("SW_ACCESS_KEY")
    config.secretKey = os.Getenv("SW_SECRET_KEY")

    return &config, nil
}
