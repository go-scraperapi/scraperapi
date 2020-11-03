package scraperapi

import (
	"context"
	"net/http"
)

// SetQueryParam adds a query parameter to the HTTP request.
func SetQueryParam(req *http.Request, key, value string) *http.Request {
	reqNew := req.Clone(context.Background())

	q := reqNew.URL.Query()
	q.Set(key, value)
	reqNew.URL.RawQuery = q.Encode()

	return reqNew
}
