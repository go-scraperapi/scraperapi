package scraperapi

import (
	"context"
	"net/http"
	"strconv"
)

type option func(r *http.Request) *http.Request

// WithContext adds context to the request.
func WithContext(ctx context.Context) option {
	return func(req *http.Request) *http.Request {
		return req.WithContext(ctx)
	}
}

// WithRenderJS makes the request fetch pages using a headless browser.
func WithRenderJS() option {
	return func(req *http.Request) *http.Request {
		return SetQueryParam(req, "render", "true")
	}
}

// WithHeader instructs Scraper API to pass a provided header.
// TODO: Refactoring to only accept a single header.
func WithHeader(key, value string) option {
	return func(req *http.Request) *http.Request {
		req.Header.Set(key, value)
		return SetQueryParam(req, "keep_headers", "true")
	}
}

// WithSessionNumber attaches a session number to a request.
func WithSessionNumber(n int) option {
	return func(req *http.Request) *http.Request {
		return SetQueryParam(req, "session_number", strconv.Itoa(n))
	}
}

const (
	CountryCodeUS        = "us"
	CountryCodeCanada    = "ca"
	CountryCodeUK        = "uk"
	CountryCodeGermany   = "de"
	CountryCodeFrance    = "fr"
	CountryCodeSpain     = "es"
	CountryCodeBrazil    = "br"
	CountryCodeMexico    = "mx"
	CountryCodeIndia     = "in"
	CountryCodeJapan     = "jp"
	CountryCodeChina     = "cn"
	CountryCodeAustralia = "au"
)

// WithCountryCode ensures your requests come from the specified location.
func WithCountryCode(countyCode string) option {
	return func(req *http.Request) *http.Request {
		return SetQueryParam(req, "country_code", countyCode)
	}
}

// WithDeviceTypeMobile passes device_type=mobile to the API.
func WithDeviceTypeMobile() option {
	return func(req *http.Request) *http.Request {
		return SetQueryParam(req, "device_type", "mobile")
	}
}
