package scraperapi

import "net/http"

// AddQueryParam adds a query parameter to the HTTP request.
func AddQueryParam(req *http.Request, key, value string) {
	q := req.URL.Query()
	q.Add(key, value)
	req.URL.RawQuery = q.Encode()
}
