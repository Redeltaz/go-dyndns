package providers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

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

type Set struct {
    ID string `json:"id"`
    Records []DNSRecord `json:"records"`
}

type Change struct {
    Set Set `json:"set"`
}

type UpdateBody struct {
    Changes []Change `json:"changes"`
}

func SwGetIP(config *conf.Config) (*DNSRecord, *error) {
    url := "https://api.scaleway.com/domain/v2beta1/dns-zones/"+config.DomainName+"/records"

    var dnsRecord DNSRecord

    client := &http.Client{}

    request, error := http.NewRequest("GET", url, nil)
    if error != nil {
        return &dnsRecord, &error
    }

    request.Header.Add("X-Auth-Token", config.SecretKey)
    response, error := client.Do(request)
    if error != nil {
        return &dnsRecord, &error
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        error := errors.New("Unknown error during record get IP request, status code : " +strconv.Itoa(response.StatusCode))
        return &dnsRecord, &error
    }

    var dnsResponse DNSResponse
    error = json.NewDecoder(response.Body).Decode(&dnsResponse)

    if error != nil {
        return &dnsRecord, &error
    }

    for _, record := range dnsResponse.Records {
        if record.Name == config.SubdomainName {
            dnsRecord = record
        }
    }

    if dnsRecord == (DNSRecord{}) {
        error := errors.New("No IP found with the domain '"+config.DomainName+"' and subdomain '"+config.SubdomainName+"'")
        return &dnsRecord, &error
    } else {
        return &dnsRecord, nil
    }
}

func SwSetIP(config *conf.Config, publicIP *string, record *DNSRecord) *error {
    url := "https://api.scaleway.com/domain/v2beta1/dns-zones/"+config.DomainName+"/records"

    body := UpdateBody{
        Changes: []Change{
            {
                Set: Set{
                    ID: record.ID,
                    Records: []DNSRecord{
                        {
                            ID: record.ID,
                            Data: *publicIP,
                            Name: record.Name,
                            Priority: record.Priority,
                            TTL: record.TTL,
                            Type: record.Type,
                            Comment: record.Comment,
                        },
                    },
                },
            },
        },
    }

    jsonBody, _ := json.Marshal(body)

    client := &http.Client{}

    request, error := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonBody))
    if error != nil {
        return &error
    }

    request.Header.Add("X-Auth-Token", config.SecretKey)
    request.Header.Set("Content-Type", "application/json")
    response, error := client.Do(request)
    if error != nil {
        return &error
    }
    defer response.Body.Close()

    if response.StatusCode == http.StatusOK {
        return nil
    } else {
        error := errors.New("Unknown error during record update IP request, status code : " + strconv.Itoa(response.StatusCode))
        return &error
    }
}
