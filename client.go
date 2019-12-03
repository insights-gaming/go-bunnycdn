package bunnycdn

import (
	"net/http"
)

type Client struct {
	hc *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		hc: &http.Client{
			Transport: &transport{
				base: http.DefaultTransport,
				key:  apiKey,
			},
		},
	}
}

func (c *Client) Zone(name string) *StorageZone {
	return &StorageZone{
		c:    c,
		name: name,
	}
}
