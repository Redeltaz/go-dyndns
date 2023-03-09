package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const publicRequestUrl string = "https://pi.ipify.org"

func main() {
    publicIP, error := getPublicIP()
    if error != nil {
        log.Fatal(*error)
        os.Exit(1)
    }

    fmt.Print(*publicIP)
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
