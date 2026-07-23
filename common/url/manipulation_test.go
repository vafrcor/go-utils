package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseURLHostPort(t *testing.T) {
	tests := map[string]map[string]string{
		"https://127.0.0.1:8080": {
			"host": "127.0.0.1",
			"port": "8080",
		},
		"127.0.0.1:8080": {
			"host": "127.0.0.1",
			"port": "8080",
		},
		"localhost:3000": {
			"host": "localhost",
			"port": "3000",
		},
		"[::1]:9090": {
			"host": "::1",
			"port": "9090",
		},
		"https://example.com": {
			"host": "example.com",
			"port": "",
		},
	}

	for u, d := range tests {
		host, port, err := ParseURLHostPort(u)
		assert.NoError(t, err)
		assert.Equal(t, d["host"], host)
		assert.Equal(t, d["port"], port)
	}
}
