# scraperapi

Scraper API SDK. See https://www.scraperapi.com/documentation for more details.

In order to use the library, you'll have to sign up and obtain an API key:
https://www.scraperapi.com/signup

## Installation

`go get -u github.com/go-scraperapi/scraperapi`

## Usage

### API Demo

```go
s := scraperapi.New("<your-api-key>")

// Scrape a URL.
s.Get("http://httpbin.org/anything")

// Pass a context to control execution.
ctx, _ := context.WithCancel(context.Background())
s.Get(
	"http://httpbin.org/anything",
	scraperapi.WithContext(ctx),
)

// Pass other options.
s.Get(
	"http://httpbin.org/anything",
	scraperapi.WithRenderJS(),
	scraperapi.WithHeader("X-MyHeader", "123"),
	scraperapi.WithSessionNumber("63c38f8dd07491e16e4d125983800a29"),
	scraperapi.WithCountryCode(scraperapi.CountryCodeAustralia),
	scraperapi.WithDeviceTypeMobile(),
	scraperapi.WithAutoParse(),
	scraperapi.WithPremium(),
)

// Scrape a form submit result.
form := url.Values{}
form.Add("foo", "bar")
s.Post(
	"http://httpbin.org/anything",
	strings.NewReader(form.Encode()),
	scraperapi.WithHeader("Content-Type", "application/x-www-form-urlencoded"),
)

// At last, request your account info and limits.
s.Account()
```

### Runnable Example 

```go
package main

import (
	"fmt"
	"github.com/go-scraperapi/scraperapi"
	"net/http/httputil"
)

func main() {
	s := scraperapi.New("<your-api-key>")

	resp, err := s.Get("http://httpbin.org/anything")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respText, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(respText))
}
```

Result:

```
HTTP/1.1 200 OK
Transfer-Encoding: chunked
Access-Control-Allow-Credentials: true
Access-Control-Allow-Headers: Origin, X-Requested-With, Content-Type, Accept
Access-Control-Allow-Methods: HEAD,GET,POST,DELETE,OPTIONS,PUT
Access-Control-Allow-Origin: undefined
Connection: keep-alive
Content-Type: application/json; charset=utf-8
Date: Tue, 03 Nov 2020 13:32:06 GMT
Etag: W/"282-6SGi09EYZihTf2gfSuZWaRlMhy0"
Sa-Final-Url: http://httpbin.org/anything
Sa-Statuscode: 200
Vary: Accept-Encoding
X-Powered-By: Express

282
{
  "args": {}, 
  "data": "", 
  "files": {}, 
  "form": {}, 
  "headers": {
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8", 
    "Accept-Encoding": "gzip, deflate", 
    "Accept-Language": "en-US,en;q=0.9,es;q=0.8", 
    "Host": "httpbin.org", 
    "Upgrade-Insecure-Requests": "1", 
    "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/CB9B9E", 
    "X-Amzn-Trace-Id": "Root=1-5fa15bd6-0c9bb6930d967d9565343e28"
  }, 
  "json": null, 
  "method": "GET", 
  "origin": "82.7.122.146", 
  "url": "http://httpbin.org/anything"
}

0
```

## Test

Unit tests:
```bash
go test -v`
```

Integration tests:

```bash
SCRAPER_API_KEY=<your-api-key> go test -v -tags=integration
```

## License

MIT
