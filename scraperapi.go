package scraperapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// BaseURL is a default base URL used to request the API.
const BaseURL = "http://api.scraperapi.com"

// Client is a main SDK struct which facilitates access to the API.
type Client struct {
	HTTPClient *http.Client
	BaseURL    string

	apiKey string
}

// New creates a new Scraper API client.
func New(apiKey string) *Client {
	return &Client{
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
		BaseURL: BaseURL,
		apiKey:  apiKey,
	}
}

// Get performs a GET HTTP request to Scraper API.
func (c *Client) Get(url string, options ...option) (*http.Response, error) {
	return c.makeAPICall("GET", url, nil, options)
}

// Post performs a POST HTTP request to Scraper API.
func (c *Client) Post(url string, body io.Reader, options ...option) (*http.Response, error) {
	return c.makeAPICall("POST", url, body, options)
}

// Put performs a PUT HTTP request to Scraper API.
func (c *Client) Put(url string, body io.Reader, options ...option) (*http.Response, error) {
	return c.makeAPICall("PUT", url, body, options)
}

func (c *Client) makeAPICall(httpMethod, url string, body io.Reader, options []option) (*http.Response, error) {
	req, err := http.NewRequest(httpMethod, c.BaseURL, body)
	if err != nil {
		return nil, fmt.Errorf("can't create an HTTP request: %s", err)
	}

	req = SetQueryParam(req, "url", url)
	for k := range options {
		req = options[k](req)
	}

	return c.sendRequest(req)
}

// AccountResponse is a response from the account API call.
type AccountResponse struct {
	ConcurrentRequests int `json:"concurrentRequests"`
	RequestCount       int `json:"requestCount"`
	FailedRequestCount int `json:"failedRequestCount"`
	RequestLimit       int `json:"requestLimit"`
	ConcurrencyLimit   int `json:"concurrencyLimit"`
}

// Account retrieves account usage information.
func (c *Client) Account() (accountResp AccountResponse, err error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/account", nil)
	if err != nil {
		return
	}

	//req = req.WithContext(ctx)

	res, err := c.sendRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&accountResp)
	return
}

func (c *Client) sendRequest(req *http.Request) (*http.Response, error) {
	req = SetQueryParam(req, "api_key", c.apiKey)
	req = SetQueryParam(req, "scraper_sdk", "go")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	// TODO: Properly handle errors.
	//if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
	//	var errRes errorResponse
	//	if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
	//		return errors.New(errRes.Message)
	//	}
	//
	//	return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	//}
	//
	//fullResponse := successResponse{
	//	Data: v,
	//}
	//if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
	//	return err
	//}

	return res, nil
}
