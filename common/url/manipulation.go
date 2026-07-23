package url

import (
	"net/url"
	"strings"
)

func ParseURLHostPort(raw string) (host, port string, err error) {
	// If scheme is missing, prepend a dummy one
	if !strings.Contains(raw, "://") {
		raw = "http://" + raw
	}

	u, err := url.Parse(raw)
	if err != nil {
		return "", "", err
	}

	host = u.Hostname()
	port = u.Port()

	return host, port, nil
}
