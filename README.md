# scraperapi

Scraper API SDK. See https://www.scraperapi.com/documentation for more details.

In order to use the library, you'll have to sign up and obtain an API key:
https://www.scraperapi.com/signup

## Installation

`go get github.com/go-scraperapi/scraperapi`

## Usage

```go
package main

import (
	"fmt"
	"github.com/go-scraperapi/scraperapi"
)

func main() {
	s := scraperapi.New("<your-api-key>")

	// Request your account info.
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
