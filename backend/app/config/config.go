package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config holds all configuration for the application
type Config struct {
	Server    ServerConfig
	Location  LocationConfig
	Cache     CacheConfig
	RateLimit RateLimitConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port string
	Host string
	Mode string
}

// LocationConfig holds location service configuration
type LocationConfig struct {
	APIURL  string
	Timeout int
	Retries int
}

// CacheConfig holds cache-related configuration
type CacheConfig struct {
	TTL     int
	MaxSize int
	Enabled bool
}

// RateLimitConfig holds rate limiting configuration
type RateLimitConfig struct {
	Enabled bool
	Limit   int
	Window  int
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	config := &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "7000"),
			Host: getEnv("HOST", "localhost"),
			Mode: getEnv("GIN_MODE", "release"),
		},
		Location: LocationConfig{
			APIURL:  getEnv("LOCATION_API_URL", "https://ipinfo.io/json"),
			Timeout: getEnvInt("LOCATION_TIMEOUT", 10),
			Retries: getEnvInt("LOCATION_RETRIES", 3),
		},
		Cache: CacheConfig{
			TTL:     getEnvInt("CACHE_TTL", 30),
			MaxSize: getEnvInt("CACHE_MAX_SIZE", 1000),
			Enabled: getEnvBool("CACHE_ENABLED", true),
		},
		RateLimit: RateLimitConfig{
			Enabled: getEnvBool("RATE_LIMIT_ENABLED", true),
			Limit:   getEnvInt("RATE_LIMIT_LIMIT", 100),
			Window:  getEnvInt("RATE_LIMIT_WINDOW", 60),
		},
	}

	// Validate configuration
	if err := config.Validate(); err != nil {
		panic(fmt.Sprintf("Invalid configuration: %v", err))
	}

	return config
}

// Validate validates the configuration
func (c *Config) Validate() error {
	// Validate server port
	if port, err := strconv.Atoi(c.Server.Port); err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("invalid port: %s", c.Server.Port)
	}

	// Validate location timeout
	if c.Location.Timeout < 1 || c.Location.Timeout > 60 {
		return fmt.Errorf("invalid location timeout: %d", c.Location.Timeout)
	}

	// Validate location retries
	if c.Location.Retries < 0 || c.Location.Retries > 10 {
		return fmt.Errorf("invalid location retries: %d", c.Location.Retries)
	}

	// Validate cache TTL
	if c.Cache.TTL < 1 || c.Cache.TTL > 3600 {
		return fmt.Errorf("invalid cache TTL: %d", c.Cache.TTL)
	}

	// Validate cache max size
	if c.Cache.MaxSize < 1 || c.Cache.MaxSize > 10000 {
		return fmt.Errorf("invalid cache max size: %d", c.Cache.MaxSize)
	}

	// Validate rate limit
	if c.RateLimit.Limit < 1 || c.RateLimit.Limit > 10000 {
		return fmt.Errorf("invalid rate limit: %d", c.RateLimit.Limit)
	}

	if c.RateLimit.Window < 1 || c.RateLimit.Window > 3600 {
		return fmt.Errorf("invalid rate limit window: %d", c.RateLimit.Window)
	}

	return nil
}

// GetServerAddress returns the full server address
func (c *Config) GetServerAddress() string {
	return fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port)
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvInt gets an environment variable as integer or returns a default value
func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getEnvBool gets an environment variable as boolean or returns a default value
func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		value = strings.ToLower(value)
		return value == "true" || value == "1" || value == "yes"
	}
	return defaultValue
}
