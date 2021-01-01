package duckdns

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"net"
	"strings"

	"github.com/jetstack/cert-manager/pkg/issuer/acme/dns/util"
)

const (
	DuckDnsBaseUrl = "https://www.duckdns.org/update?"
)

type Client struct {
	httpClient http.Client
	apiToken   string
}

func newClient(apiToken string) (*Client, error) {
	return &Client{ 
		httpClient: http.Client{ Timeout: 30 * time.Second, }, 
		apiToken: apiToken,
		}, nil
}

func (c *Client) RecordsUrl(entry string) string {
	return fmt.Sprintf("%sdomains=%s&token=%s", DuckDnsBaseUrl, entry, c.apiToken)
}

func (c *Client) getDomainAndEntry(zone string, fqdn string) (string, string) {
	// Both ch.ResolvedZone and ch.ResolvedFQDN end with a dot: '.'
	entry := util.UnFqdn(strings.TrimSuffix(fqdn, zone))
	entry = util.UnFqdn(entry)
	domain := util.UnFqdn(zone)
	return entry, domain
}

func (c *Client) addTxtRecord(entry, key string) error {
	// curl https://www.duckdns.org/update?domains=<DOMAIN>&token=<TOKEN>&txt=<TXT>

	url := fmt.Sprintf("%s&txt=%s", c.RecordsUrl(entry), key)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("An error occured in the request, Status Code %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Unable to read http response", err)
	}

	//Check if response contains OK or KO
	if !strings.HasPrefix(string(body), "OK") {
		return fmt.Errorf("Response was not OK, Resp %v", body)
	}

	return nil
}

func (c *Client) getTxtRecord(entry string) (string, error) {

	txt, err := net.LookupTXT(entry)
	if err != nil {
		return "", fmt.Errorf("Unable to get txt record", err)
	}

	if len(txt) == 0 {
		return "", nil
	}

	//duckdns should have only 1 record
	return txt[0], nil
}

func (c *Client) deleteTxtRecord(entry, key string) error {
	// curl https://www.duckdns.org/update?domains=<DOMAIN>&token=<TOKEN>&clear=true

	url := fmt.Sprintf("%s&txt=%s&clear=true", c.RecordsUrl(entry), key)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("An error occured in the request, Status Code %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Unable to read http response", err)
	}

	//Check if response contains OK or KO
	if !strings.HasPrefix(string(body), "OK") {
		return fmt.Errorf("Response was not OK, Resp %v", body)
	}

	return nil
}