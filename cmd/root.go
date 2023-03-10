package cmd

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Redeltaz/go-dyndns/conf"
	"github.com/Redeltaz/go-dyndns/providers"
)

const publicRequestUrl string = "https://api.ipify.org"

func Root() {
	publicIP, error := getPublicIP()
	if error != nil {
		log.Fatal(*error)
	}

	config, error := conf.Loadenv()
	if error != nil {
		log.Fatal(*error)
	}

	record, error := providers.SwGetIP(config)
	if error != nil {
		log.Fatal(*error)
	}

	if record.Data != *publicIP {
		error := providers.SwSetIP(config, publicIP, record)
		if error != nil {
			log.Fatal(error)
		} else {
			log.Fatal("Public IP successfully changed !")
		}
	} else {
		log.Fatal("Public IP didn't change, no need to update")
	}
}

func getPublicIP() (*string, *error) {
	var body string

	response, error := http.Get(publicRequestUrl)
	if error != nil {
		return &body, &error
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, error := io.ReadAll(response.Body)
		if error != nil {
			return &body, &error
		}

		body = string(bodyBytes)

		return &body, nil
	} else {
		error := errors.New("Unknown error during public IP request, status code : " + strconv.Itoa(response.StatusCode))
		return &body, &error
	}
}
