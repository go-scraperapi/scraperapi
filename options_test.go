package scraperapi

import (
	"context"
	"net/http"
	"testing"
)

func TestWithContext(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://google.com", nil)

	req = WithContext(context.WithValue(context.Background(), "foo", "bar"))(req)

	actual := req.Context().Value("foo")
	if actual != "bar" {
		t.Errorf("expected %s, got %v", "bar", actual)
	}
}

func TestWithRenderJS(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://google.com", nil)

	req = WithRenderJS()(req)

	expected := "true"
	actual := req.URL.Query().Get("render")
	if actual != expected {
		t.Errorf("expected %s, got %v", expected, actual)
	}
}

func TestWithHeaders(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://google.com", nil)

	expected := "application/json"

	req = WithHeader("Content-Type", expected)(req)

	actual := req.Header.Get("Content-Type")
	if actual != expected {
		t.Errorf("expected %s, got %v", expected, actual)
	}

	actual2 := req.URL.Query().Get("keep_headers")
	if actual2 != "true" {
		t.Errorf("expected %s, got %v", "true", actual2)
	}
}

func TestWithSessionNumber(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://google.com", nil)

	req = WithSessionNumber("150")(req)

	expected := "150"
	actual := req.URL.Query().Get("session_number")
	if actual != expected {
		t.Errorf("expected %s, got %v", expected, actual)
	}
}

func TestWithCountryCode(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://google.com", nil)

	req = WithCountryCode(CountryCodeUS)(req)

	expected := "us"
	actual := req.URL.Query().Get("country_code")
	if actual != expected {
		t.Errorf("expected %s, got %v", expected, actual)
	}
}

func TestWithDeviceTypeMobile(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://google.com", nil)

	req = WithDeviceType(DeviceTypeMobile)(req)

	expected := "mobile"
	actual := req.URL.Query().Get("device_type")
	if actual != expected {
		t.Errorf("expected %s, got %v", expected, actual)
	}
}

func TestWithAutoParse(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://google.com", nil)

	req = WithAutoParse()(req)

	expected := "true"
	actual := req.URL.Query().Get("autoparse")
	if actual != expected {
		t.Errorf("expected %s, got %v", expected, actual)
	}
}

func TestWithPremium(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://google.com", nil)

	req = WithPremium()(req)

	expected := "true"
	actual := req.URL.Query().Get("premium")
	if actual != expected {
		t.Errorf("expected %s, got %v", expected, actual)
	}
}
