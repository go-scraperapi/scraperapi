package scraperapi

// Client is a main SDK struct which facilitates access to the API.
type Client struct {
}

// New creates a new Scraper API client.
func New() *Client {
	return &Client{}
}
