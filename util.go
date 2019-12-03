package bunnycdn

import (
	"net/url"
	"strings"
)

func ParseURI(uri string) (zone, path, name string, _ error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", "", "", nil
	}

	if u.Scheme != "b-cdn" {
		return "", "", "", ErrInvalidURI
	}

	zone = u.Host

	if pos := strings.LastIndex(u.Path, "/"); pos != -1 {
		path = u.Path[:pos]
		name = u.Path[pos+1:]
	} else {
		name = u.Path
	}

	return
}
