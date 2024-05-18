package steamclient

import (
	"net/http"
	"testing"
)

func TestNewClientWithoutId(t *testing.T) {
	htCl := http.Client{}

	client := NewClientWithoutKey(&htCl)
	if client.Key != "" {
		t.Errorf("Expected empty key, got %v", client.Key)
	}

	if client.IsKeySet() {
		t.Errorf("Expected IsKeySet to return false for empty key")
	}
}

func TestNew(t *testing.T) {
	testKey := "my-api-key"
	htCl := http.Client{}
	client := New(testKey, &htCl)
	if client.Key != testKey {
		t.Errorf("Expected key %v, got %v", testKey, client.Key)
	}

	if !client.IsKeySet() {
		t.Errorf("Expected IsKeySet to return true for non-empty key")
	}
}
