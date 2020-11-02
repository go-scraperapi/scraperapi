// +build integration

package scraperapi

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestClient_Account(t *testing.T) {
	c := New(os.Getenv("SCRAPER_API_KEY"))

	_, err := c.Account()
	if err != nil {
		t.Error(err)
	}
}

func TestClient_Get(t *testing.T) {
	c := New(os.Getenv("SCRAPER_API_KEY"))

	resp, err := c.Get("http://httpbin.org/ip")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	respText, _ := ioutil.ReadAll(resp.Body)
	if string(respText) == "" {
		t.Error("resp is empty")
	}
	t.Log(string(respText))
}
