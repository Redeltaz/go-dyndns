package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	//"fmt"
	"io"
	//"log"
	"net/http"
	//"os"
	"github.com/Redeltaz/go-dyndns/conf"
	"github.com/Redeltaz/go-dyndns/providers"
)

const publicRequestUrl string = "https://api.ipify.org"

func Root() {
    publicIP, error := getPublicIP()
    if error != nil {
        log.Fatal(*error)
        os.Exit(1)
    }

    config, error := conf.Loadenv()
    if error != nil {
        log.Fatal(*error)
        os.Exit(1)
    }

    recordIP, error := providers.SwGetIP(config)
    if error != nil {
        log.Fatal(*error)
        os.Exit(1)
    }

    fmt.Println(*recordIP == *publicIP)
}

func getPublicIP() (*string, *error) {
    var body string

    res, error := http.Get(publicRequestUrl)
    if error != nil {
       return &body, &error
    }
    defer res.Body.Close()

    if res.StatusCode == http.StatusOK {
        bodyBytes, error := io.ReadAll(res.Body)
        if error != nil {
            return &body, &error
        }

        body = string(bodyBytes)

        return &body, nil
    } else {
        error := errors.New("Status code is not ok")
        return &body, &error
    }
}
