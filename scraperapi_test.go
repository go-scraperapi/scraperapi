package scraperapi

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestClient_Post(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))

		h1Expected := "application/x-www-form-urlencoded"
		h1Actual := r.Header.Get("Content-Type")

		h2Expected := "special value"
		h2Actual := r.Header.Get("X-My-Header")

		apiKeyExpected := "very-secret-api-key"
		apiKeyActual := r.URL.Query().Get("api_key")

		sdkExpected := "go"
		sdkActual := r.URL.Query().Get("scraper_sdk")

		r.ParseForm()
		fExpected := "bar"
		fActual := r.Form["foo"][0]

		if h1Expected != h1Actual {
			t.Errorf("expected Content-Type %s, got %s", h1Expected, h1Actual)
		}

		if h2Expected != h2Actual {
			t.Errorf("expected X-My-Header %s, got %s", h2Expected, h2Actual)
		}

		if apiKeyExpected != apiKeyActual {
			t.Errorf("expected api_key %s, got %s", apiKeyExpected, apiKeyActual)
		}

		if sdkExpected != sdkActual {
			t.Errorf("expected scraper_sdk %s, got %s", sdkExpected, sdkActual)
		}

		if fExpected != fActual {
			t.Errorf("expected form value %s, got %s", fExpected, fActual)
		}
	}))
	defer srv.Close()

	c := New("very-secret-api-key")
	c.HTTPClient = srv.Client()
	c.BaseURL = srv.URL

	form := url.Values{}
	form.Add("foo", "bar")
	_, err := c.Post(
		"http://httpbin.org/anything",
		strings.NewReader(form.Encode()),
		WithHeader("Content-Type", "application/x-www-form-urlencoded"),
		WithHeader("X-My-Header", "special value"),
	)
	if err != nil {
		t.Error(err)
	}
}
