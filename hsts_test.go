package hsts

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestAddHSTSHeader(t *testing.T) {
	// Create mock HTTP handler to test
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		// Do nothing
	}

	// Wrap mock handler with AddHSTSHeader middleware
	wrappedHandlerFunc := AddHSTSHeader(http.HandlerFunc(handlerFunc))

	// Create mock http request
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create mock http response recorder
	rr := httptest.NewRecorder()

	// Call wrapped handler with mock request and response recorder
	wrappedHandlerFunc(rr, req)

	// Check if Strict-Transport-Security header added to the response
	expectedHeaderValue := "max-age=" + strconv.FormatInt(int64(MaxAge/time.Second), 10)
	if IncludeSubdomains {
		expectedHeaderValue += "; includeSubDomains"
	}
	if Preload {
		expectedHeaderValue += "; preload"
	}
	if rr.Header().Get("Strict-Transport-Security") != expectedHeaderValue {
		t.Errorf("Expected header value '%s', but got '%s'", expectedHeaderValue, rr.Header().Get("Strict-Transport-Security"))
	}
}
