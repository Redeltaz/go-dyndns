package providers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Redeltaz/go-dyndns/conf"
)

type DNSRecord struct {
    ID       string `json:"id"`
    Data     string `json:"data"`
    Name     string `json:"name"`
    Priority int    `json:"priority"`
    TTL      int    `json:"ttl"`
    Type     string `json:"type"`
    Comment  string `json:"comment"`
}

type DNSResponse struct {
    Records    []DNSRecord `json:"records"`
}

func SwGetIP(config *conf.Config) (*string, *error) {
    url := "https://api.scaleway.com/domain/v2beta1/dns-zones/"+config.DomainName+"/records"

    var recordIP string

    client := &http.Client{}

    request, error := http.NewRequest("GET", url, nil)
    if error != nil {
        return &recordIP, &error
    }

    request.Header.Add("X-Auth-Token", config.SecretKey)
    response, error := client.Do(request)
    if error != nil {
        return &recordIP, &error
    }
    defer response.Body.Close()

    var dnsResponse DNSResponse
    error = json.NewDecoder(response.Body).Decode(&dnsResponse)

    if error != nil {
        return &recordIP, &error
    }

    for _, record := range dnsResponse.Records {
        if record.Name == config.SubdomainName {
            recordIP = record.Data
        }
    }

    if recordIP == "" {
        error := errors.New("No IP found with the domain "+config.DomainName+" and subdomain "+config.SubdomainName)
        return &recordIP, &error
    } else {
        return &recordIP, nil
    }
}
