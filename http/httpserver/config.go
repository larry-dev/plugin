package httpserver

import (
	"net/http"
)

type Config struct {
	ListenAddr  string
	Handler     http.Handler
	EnableHttps bool
	// Server requires TLS client certificate authentication
	CertFile string
	// Server requires TLS client certificate authentication
	KeyFile string
	// the password to decrypt the certificate
	Password string
}

func DefaultConfig() Config {
	return Config{
		ListenAddr:  ":8080",
	}
}
