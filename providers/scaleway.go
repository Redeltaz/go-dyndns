package providers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Redeltaz/go-dyndns/conf"
)

func SwGetIP(config *conf.Config) (*string, *error) {
    url := "https://api.scaleway.com/domain/v2beta1/dns-zones/"+config.DomainName+"/records"

    var body string

    client := &http.Client{}

    request, error := http.NewRequest("GET", url, nil)
    if error != nil {
        return &body, &error
    }
    request.Header.Add("X-Auth-Token", config.SecretKey)
    response, error := client.Do(request)
    if error != nil {
        return &body, &error
    }
    defer response.Body.Close()

    bodyBytes, error := io.ReadAll(response.Body)
    fmt.Println(string(bodyBytes))

    return &body, nil
}
