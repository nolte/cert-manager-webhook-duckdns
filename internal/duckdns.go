package internal

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	DuckDnsBaseUrl = "https://www.duckdns.org/update?"
)

type DuckDNSClient struct {
	apiKey              string
}

func NewDuckDNSClient(apiKey string) *DuckDNSClient {
	return &DuckDNSClient{
		apiKey: apiKey,
	}
}

func (c *DuckDNSClient) RecordsUrl(domain string) string {
	return fmt.Sprintf("%sdomains=%s&token=%s", DuckDnsBaseUrl, domain, c.apiKey)
}

func (c *DuckDNSClient) doRequest(req *http.Request, readResponseBody bool) (int, []byte, error) {
	
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}

	if res.StatusCode == http.StatusOK && readResponseBody {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return 0, nil, err
		}
		return res.StatusCode, data, nil
	}

	return res.StatusCode, nil, nil
}

func (c *DuckDNSClient) CreateTxtRecord(domain *string, value *string) error {
	
	// curl https://www.duckdns.org/update?domains=<DOMAIN>&token=<TOKEN>&txt=<TXT>

	url := fmt.Sprintf("%s&txt=%s", c.RecordsUrl(*domain), *value)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	status, _, err := c.doRequest(req, false)
	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return fmt.Errorf("failed creating TXT record: %v", err)
	}

	return nil
}

func (c *DuckDNSClient) DeleteTxtRecord(domain *string, value *string) error {
	
	// curl https://www.duckdns.org/update?domains=<DOMAIN>&token=<TOKEN>&txt=<TXT>&clear=true

	url := fmt.Sprintf("%s&txt=%s&clear=true", c.RecordsUrl(*domain), *value)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	status, _, err := c.doRequest(req, false)
	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return fmt.Errorf("failed creating TXT record: %v", err)
	}

	return nil
}
