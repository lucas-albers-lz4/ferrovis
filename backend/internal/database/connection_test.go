package database

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name     string
		envVars  map[string]string
		expected *Config
	}{
		{
			name:    "default configuration",
			envVars: map[string]string{},
			expected: &Config{
				Host:     "localhost",
				Port:     "5432",
				User:     "postgres",
				Password: "ferrovis123",
				DBName:   "ferrovis_dev",
				SSLMode:  "disable",
			},
		},
		{
			name: "custom configuration",
			envVars: map[string]string{
				"DB_HOST":     "custom-host",
				"DB_PORT":     "3306",
				"DB_USER":     "custom-user",
				"DB_PASSWORD": "secret",
				"DB_NAME":     "custom_db",
				"DB_SSLMODE":  "require",
			},
			expected: &Config{
				Host:     "custom-host",
				Port:     "3306",
				User:     "custom-user",
				Password: "secret",
				DBName:   "custom_db",
				SSLMode:  "require",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables
			for key, value := range tt.envVars {
				t.Setenv(key, value)
			}

			// Load configuration
			config := LoadConfig()

			// Verify configuration
			if config.Host != tt.expected.Host {
				t.Errorf("Expected Host %s, got %s", tt.expected.Host, config.Host)
			}
			if config.Port != tt.expected.Port {
				t.Errorf("Expected Port %s, got %s", tt.expected.Port, config.Port)
			}
			if config.User != tt.expected.User {
				t.Errorf("Expected User %s, got %s", tt.expected.User, config.User)
			}
			if config.Password != tt.expected.Password {
				t.Errorf("Expected Password %s, got %s", tt.expected.Password, config.Password)
			}
			if config.DBName != tt.expected.DBName {
				t.Errorf("Expected DBName %s, got %s", tt.expected.DBName, config.DBName)
			}
			if config.SSLMode != tt.expected.SSLMode {
				t.Errorf("Expected SSLMode %s, got %s", tt.expected.SSLMode, config.SSLMode)
			}
		})
	}
}

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		defaultValue string
		envValue     string
		expected     string
	}{
		{
			name:         "environment variable exists",
			key:          "TEST_VAR",
			defaultValue: "default",
			envValue:     "from_env",
			expected:     "from_env",
		},
		{
			name:         "environment variable does not exist",
			key:          "NONEXISTENT_VAR",
			defaultValue: "default",
			envValue:     "",
			expected:     "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clean up environment variable
			defer func() {
				if err := os.Unsetenv(tt.key); err != nil {
					t.Logf("Warning: failed to unset environment variable %s: %v", tt.key, err)
				}
			}()

			// Set environment variable if needed
			if tt.envValue != "" {
				if err := os.Setenv(tt.key, tt.envValue); err != nil {
					t.Fatalf("Failed to set environment variable %s: %v", tt.key, err)
				}
			}

			result := getEnv(tt.key, tt.defaultValue)

			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
