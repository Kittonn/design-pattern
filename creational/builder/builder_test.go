package builder

import (
	"testing"
)

func TestServerConfigBuilder(t *testing.T) {
	builder := NewServerConfigBuilder()

	config, err := builder.SetHost("localhost").
		SetPort(8080).
		EnableTLS().
		SetTimeout(60).
		SetMaxRetries(5).
		Build()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if config.Host != "localhost" {
		t.Errorf("Expected Host to be 'localhost', got '%s'", config.Host)
	}

	if config.Port != 8080 {
		t.Errorf("Expected Port to be 8080, got %d", config.Port)
	}

	if !config.TLS {
		t.Error("Expected TLS to be enabled")
	}

	if config.Timeout != 60 {
		t.Errorf("Expected Timeout to be 60, got %d", config.Timeout)
	}

	if config.MaxRetries != 5 {
		t.Errorf("Expected MaxRetries to be 5, got %d", config.MaxRetries)
	}
}