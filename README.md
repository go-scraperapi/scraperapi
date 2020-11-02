# scraperapi

Scraper API SDK. See https://www.scraperapi.com/documentation for more details.

In order to use the library, you'll have to sign up and obtain an API key:
https://www.scraperapi.com/signup

## Installation

`go get -u github.com/go-scraperapi/scraperapi`

## Example Usage

```go
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"github.com/go-scraperapi/scraperapi"
)

func main() {
	s := scraperapi.New("<your-api-key>")

	// Scrape a URL.
	resp, err := s.Get("http://httpbin.org/ip")

	// Pass a context to control execution.
	ctx, _ := context.WithCancel(context.Background())
	resp, err = s.Get(
		"http://httpbin.org/ip",
		scraperapi.WithContext(ctx),
	)

	// You can also pass a number of other options.
	resp, err = s.Get(
		"http://httpbin.org/ip",
		scraperapi.WithRenderJS(),
		scraperapi.WithHeaders(map[string]string{"X-MyHeader": "123"}),
		scraperapi.WithSessionNumber(45),
		scraperapi.WithCountryCode(scraperapi.CountryCodeAustralia),
	)

	// For the sake of the example, let's print it out.
	if err != nil {
		fmt.Println("error making a request: %v", err)
		return
	}
	defer resp.Body.Close()

	respText, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respText), err)

	// At last, request your account info and limits.
	accountInfo, err := s.Account()
	fmt.Printf("%+v %v", accountInfo, err)
}
``` 

## Test

Run integration tests making real HTTP requests:

```bash
SCRAPER_API_KEY=<your-api-key> go test -v -tags=integration
```

## License

MIT
