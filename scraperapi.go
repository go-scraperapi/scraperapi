package scraperapi

import (
	"encoding/json"
	"fmt"
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

func (c *Client) Get(url string) (*http.Response, error) {
	return c.makeAPICall("GET", url)
}

func (c *Client) makeAPICall(httpMethod, url string) (*http.Response, error) {
	req, err := http.NewRequest(httpMethod, c.BaseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("can't create an HTTP request: %s", err)
	}

	AddQueryParam(req, "url", url)

	return c.sendRequest(req)

	//req = req.WithContext(ctx)
}

// TODO: Implement, add POST, PUT

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

// TODO: applyOptions

func (c *Client) sendRequest(req *http.Request) (*http.Response, error) {
	AddQueryParam(req, "api_key", c.apiKey)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

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
