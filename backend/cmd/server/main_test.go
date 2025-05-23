package main

import (
	"net/http"
	"testing"
)

func TestConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant interface{}
		expected interface{}
	}{
		{"statusOK", statusOK, http.StatusOK},
		{"statusInternalServerError", statusInternalServerError, http.StatusInternalServerError},
		{"defaultTableCount", defaultTableCount, 10},
		{"weekDuration", weekDuration, 12},
		{"defaultPort", defaultPort, "8080"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("Expected %s to be %v, got %v", tt.name, tt.expected, tt.constant)
			}
		})
	}
}

func TestHTTPStatusCodes(t *testing.T) {
	// Verify our constants match standard HTTP status codes
	if statusOK != 200 {
		t.Errorf("Expected statusOK to be 200, got %d", statusOK)
	}

	if statusInternalServerError != 500 {
		t.Errorf("Expected statusInternalServerError to be 500, got %d", statusInternalServerError)
	}
}

func TestBusinessLogicConstants(t *testing.T) {
	// Test that business logic constants are reasonable
	if defaultTableCount <= 0 {
		t.Errorf("defaultTableCount should be positive, got %d", defaultTableCount)
	}

	if weekDuration <= 0 {
		t.Errorf("weekDuration should be positive, got %d", weekDuration)
	}

	if defaultPort == "" {
		t.Error("defaultPort should not be empty")
	}

	// Test that defaultPort is a valid port string
	if defaultPort != "8080" {
		t.Errorf("Expected defaultPort to be '8080', got '%s'", defaultPort)
	}
}
