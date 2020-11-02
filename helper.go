package scraperapi

import (
	"context"
	"net/http"
)

// AddQueryParam adds a query parameter to the HTTP request.
func AddQueryParam(req *http.Request, key, value string) *http.Request {
	reqNew := req.Clone(context.Background())

	q := reqNew.URL.Query()
	q.Add(key, value)
	reqNew.URL.RawQuery = q.Encode()

	return reqNew
}
