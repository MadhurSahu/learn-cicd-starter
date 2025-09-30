package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	_, err := GetAPIKey(headers)

	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Errorf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
	}

	headers.Set("Authorization", "Api 1234")
	_, err = GetAPIKey(headers)

	if !errors.Is(err, ErrMalformedAuthHeader) {
		t.Errorf("expected error %v, got %v", ErrMalformedAuthHeader, err)
	}

	headers.Set("Authorization", "ApiKey 1234")
	key, err := GetAPIKey(headers)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if key != "1234" {
		t.Errorf("expected key 1234, got %v", key)
	}
}
