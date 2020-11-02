// +build integration

package scraperapi

import (
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
