package builder

import "fmt"

type ServerConfig struct {
	Host       string
	Port       int
	TLS        bool
	Timeout    int
	MaxRetries int
}

type ServerConfigBuilder struct {
	host       string
	port       int
	tls        bool
	timeout    int
	maxRetries int
}

func NewServerConfigBuilder() *ServerConfigBuilder {
	return &ServerConfigBuilder{
		host:    "127.0.0.1",
		port:    80,
		timeout: 30,
	}
}

func (b *ServerConfigBuilder) SetHost(host string) *ServerConfigBuilder {
	b.host = host
	return b
}

func (b *ServerConfigBuilder) SetPort(port int) *ServerConfigBuilder {
	b.port = port
	return b
}

func (b *ServerConfigBuilder) EnableTLS() *ServerConfigBuilder {
	b.tls = true
	return b
}

func (b *ServerConfigBuilder) SetTimeout(timeout int) *ServerConfigBuilder {
	b.timeout = timeout
	return b
}

func (b *ServerConfigBuilder) SetMaxRetries(maxRetries int) *ServerConfigBuilder {
	b.maxRetries = maxRetries
	return b
}

func (b *ServerConfigBuilder) Build() (*ServerConfig, error) {
	if b.port <= 0 {
		return nil, fmt.Errorf("invalid port")
	}

	return &ServerConfig{
		Host:       b.host,
		Port:       b.port,
		TLS:        b.tls,
		Timeout:    b.timeout,
		MaxRetries: b.maxRetries,
	}, nil
}
