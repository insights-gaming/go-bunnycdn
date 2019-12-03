package bunnycdn

import (
	"net/http"
)

type transport struct {
	base http.RoundTripper

	key string
}

func (t *transport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("AccessKey", t.key)
	return t.base.RoundTrip(r)
}