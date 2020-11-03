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
		return AddQueryParam(req, "render", "true")
	}
}

// WithHeaders instructs Scraper API to pass provided headers.
// TODO: Refactoring to only accept a single header.
func WithHeaders(headers map[string]string) option {
	return func(req *http.Request) *http.Request {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
		return AddQueryParam(req, "keep_headers", "true")
	}
}

// WithSessionNumber attaches a session number to a request.
func WithSessionNumber(n int) option {
	return func(req *http.Request) *http.Request {
		return AddQueryParam(req, "session_number", strconv.Itoa(n))
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
		return AddQueryParam(req, "country_code", countyCode)
	}
}
