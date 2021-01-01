package duckdns

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"net"
	"strings"

	"k8s.io/klog"
)

const (
	DuckDnsBaseUrl = "https://www.duckdns.org/update?"
)

type Client struct {
	httpClient http.Client
	domain     string
	apiToken   string
}

func newClient(domain, apiToken string) (*Client, error) {
	return &Client{ 
		httpClient: http.Client{ Timeout: 30 * time.Second, },
		domain: domain,
		apiToken: apiToken,
		}, nil
}

func (c *Client) RecordsUrl(entry string) string {
	return fmt.Sprintf("%sdomains=%s&token=%s", DuckDnsBaseUrl, entry, c.apiToken)
}

func (c *Client) addTxtRecord(domain, key string) error {
	// curl https://www.duckdns.org/update?domains=<DOMAIN>&token=<TOKEN>&txt=<TXT>

	url := fmt.Sprintf("%s&txt=%s", c.RecordsUrl(domain), key)
	klog.Infof("addTxtRecord: Sending request to url: %v ", url)
	
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

func (c *Client) getTxtRecord(domain string) (string, error) {

	txt, err := net.LookupTXT(domain)
	if err != nil {
		return "", fmt.Errorf("Unable to get txt record", err)
	}

	if len(txt) == 0 {
		return "", nil
	}

	//duckdns should have only 1 record
	return txt[0], nil
}

func (c *Client) deleteTxtRecord(domain, key string) error {
	// curl https://www.duckdns.org/update?domains=<DOMAIN>&token=<TOKEN>&clear=true

	url := fmt.Sprintf("%s&txt=%s&clear=true", c.RecordsUrl(domain), key)
	klog.Infof("deleteTxtRecord: Sending request to url: %v ", url)
	
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