// File: security_config.go
//
// Responsibility:
//   - Define the configurable HTTP security policy.
//   - Load security settings from config/security.yaml.
//   - Provide safe built-in defaults when the YAML file is absent.
//   - Validate configuration before the HTTP server starts.
//
// Receives:
//   - Path to the security YAML file.
//
// Produces:
//   - Validated SecurityConfig ready for the HTTP infrastructure.
//   - Information indicating whether built-in defaults were used.
//
// Previous logical stage:
//   - cmd/web/main.go.
//
// Next logical stage:
//   - server.go and HTTP security middleware.
//
// Important restrictions:
//   - Must not contain conversion logic.
//   - Must not inspect units or scientific catalogs.
//   - Must not modify the YAML file.
//   - Unknown YAML fields must not be silently ignored.
//   - An existing but invalid YAML file must prevent server startup.
//   - Absence of the YAML file must fall back to safe defaults.
package web

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	defaultMaxBodyBytes = int64(4096)

	defaultReadHeaderTimeout = 5 * time.Second
	defaultReadTimeout       = 10 * time.Second
	defaultWriteTimeout      = 15 * time.Second
	defaultIdleTimeout       = 60 * time.Second
)

// SecurityConfig contains the validated Web security policy
// used by the HTTP infrastructure.
type SecurityConfig struct {
	HTTP     HTTPConfig
	Security SecurityOptions
	Logging  LoggingConfig
}

// HTTPConfig contains HTTP resource limits.
type HTTPConfig struct {
	MaxBodyBytes int64
	Timeouts     HTTPTimeoutConfig
}

// HTTPTimeoutConfig contains the time limits applied
// by net/http.Server.
type HTTPTimeoutConfig struct {
	ReadHeader time.Duration
	Read       time.Duration
	Write      time.Duration
	Idle       time.Duration
}

// SecurityOptions contains configurable defensive behavior.
type SecurityOptions struct {
	Headers HeadersConfig
}

// HeadersConfig controls defensive HTTP response headers.
type HeadersConfig struct {
	Enabled bool
}

// LoggingConfig controls technical HTTP request logging.
type LoggingConfig struct {
	Enabled bool
}

// SecurityConfigResult represents the result of loading
// the external security policy.
type SecurityConfigResult struct {
	Config       SecurityConfig
	UsedDefaults bool
}

// rawSecurityConfig represents the YAML document before
// duration strings are parsed and validated.
type rawSecurityConfig struct {
	HTTP struct {
		MaxBodyBytes int64 `yaml:"max_body_bytes"`

		Timeouts struct {
			ReadHeader string `yaml:"read_header"`
			Read       string `yaml:"read"`
			Write      string `yaml:"write"`
			Idle       string `yaml:"idle"`
		} `yaml:"timeouts"`
	} `yaml:"http"`

	Security struct {
		Headers struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"headers"`
	} `yaml:"security"`

	Logging struct {
		Enabled bool `yaml:"enabled"`
	} `yaml:"logging"`
}

// LoadSecurityConfig loads and validates the Web security configuration.
//
// If the file does not exist, safe built-in defaults are returned.
// If the file exists but cannot be parsed or validated, an error is returned.
func LoadSecurityConfig(
	path string,
) (SecurityConfigResult, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return SecurityConfigResult{
				Config:       defaultSecurityConfig(),
				UsedDefaults: true,
			}, nil
		}

		return SecurityConfigResult{}, fmt.Errorf(
			"leer configuración de seguridad %q: %w",
			path,
			err,
		)
	}

	raw := defaultRawSecurityConfig()

	decoder := yaml.NewDecoder(
		bytes.NewReader(data),
	)

	// A typo in a security option must not be silently ignored.
	decoder.KnownFields(true)

	if err := decoder.Decode(&raw); err != nil {
		return SecurityConfigResult{}, fmt.Errorf(
			"analizar configuración de seguridad %q: %w",
			path,
			err,
		)
	}

	config, err := validateRawSecurityConfig(raw)
	if err != nil {
		return SecurityConfigResult{}, fmt.Errorf(
			"configuración de seguridad inválida %q: %w",
			path,
			err,
		)
	}

	return SecurityConfigResult{
		Config:       config,
		UsedDefaults: false,
	}, nil
}

// defaultSecurityConfig returns the safe policy compiled
// into UConversor.
func defaultSecurityConfig() SecurityConfig {
	return SecurityConfig{
		HTTP: HTTPConfig{
			MaxBodyBytes: defaultMaxBodyBytes,
			Timeouts: HTTPTimeoutConfig{
				ReadHeader: defaultReadHeaderTimeout,
				Read:       defaultReadTimeout,
				Write:      defaultWriteTimeout,
				Idle:       defaultIdleTimeout,
			},
		},
		Security: SecurityOptions{
			Headers: HeadersConfig{
				Enabled: true,
			},
		},
		Logging: LoggingConfig{
			Enabled: true,
		},
	}
}

// defaultRawSecurityConfig prepares YAML values with safe defaults.
//
// This allows the YAML file to override only the options that
// intentionally need to change while preserving safe values for
// omitted known fields.
func defaultRawSecurityConfig() rawSecurityConfig {
	var raw rawSecurityConfig

	raw.HTTP.MaxBodyBytes = defaultMaxBodyBytes
	raw.HTTP.Timeouts.ReadHeader = defaultReadHeaderTimeout.String()
	raw.HTTP.Timeouts.Read = defaultReadTimeout.String()
	raw.HTTP.Timeouts.Write = defaultWriteTimeout.String()
	raw.HTTP.Timeouts.Idle = defaultIdleTimeout.String()

	raw.Security.Headers.Enabled = true
	raw.Logging.Enabled = true

	return raw
}

// validateRawSecurityConfig converts textual YAML values into the
// validated configuration consumed by the HTTP server.
func validateRawSecurityConfig(
	raw rawSecurityConfig,
) (SecurityConfig, error) {
	if raw.HTTP.MaxBodyBytes <= 0 {
		return SecurityConfig{}, fmt.Errorf(
			"http.max_body_bytes debe ser mayor que cero",
		)
	}

	readHeader, err := parsePositiveDuration(
		"http.timeouts.read_header",
		raw.HTTP.Timeouts.ReadHeader,
	)
	if err != nil {
		return SecurityConfig{}, err
	}

	read, err := parsePositiveDuration(
		"http.timeouts.read",
		raw.HTTP.Timeouts.Read,
	)
	if err != nil {
		return SecurityConfig{}, err
	}

	write, err := parsePositiveDuration(
		"http.timeouts.write",
		raw.HTTP.Timeouts.Write,
	)
	if err != nil {
		return SecurityConfig{}, err
	}

	idle, err := parsePositiveDuration(
		"http.timeouts.idle",
		raw.HTTP.Timeouts.Idle,
	)
	if err != nil {
		return SecurityConfig{}, err
	}

	return SecurityConfig{
		HTTP: HTTPConfig{
			MaxBodyBytes: raw.HTTP.MaxBodyBytes,
			Timeouts: HTTPTimeoutConfig{
				ReadHeader: readHeader,
				Read:       read,
				Write:      write,
				Idle:       idle,
			},
		},
		Security: SecurityOptions{
			Headers: HeadersConfig{
				Enabled: raw.Security.Headers.Enabled,
			},
		},
		Logging: LoggingConfig{
			Enabled: raw.Logging.Enabled,
		},
	}, nil
}

// parsePositiveDuration parses one timeout and guarantees
// that it represents a positive interval.
func parsePositiveDuration(
	name string,
	value string,
) (time.Duration, error) {
	duration, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf(
			"%s contiene una duración inválida %q: %w",
			name,
			value,
			err,
		)
	}

	if duration <= 0 {
		return 0, fmt.Errorf(
			"%s debe ser mayor que cero",
			name,
		)
	}

	return duration, nil
}
