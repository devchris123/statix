package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestNewServer tests the NewServer function.
func TestNewServer(t *testing.T) {
	addr := ":8080"
	srv := NewServer(addr)

	if srv.server.Addr != addr {
		t.Errorf("expected server Addr %s, got %s", addr, srv.server.Addr)
	}
}

// Test that request.URL.path is always the relative path
// regardless of whether the request is made with a full URL or a relative path.
func TestRequestPath(t *testing.T) {
	path := "/testfile.txt"

	// Test with a relative path
	r, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	if r.URL.Path != path {
		t.Errorf("expected request URL path %s, got %s", path, r.URL.Path)
	}

	// Test with a full URL
	r, err = http.NewRequest("GET", "http://example.com"+path, nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	if r.URL.Path != path {
		t.Errorf("expected request URL path %s, got %s", path, r.URL.Path)
	}
}

// TestStaticFileHandler tests the staticFileHandler function.
func TestStaticFileHandler(t *testing.T) {
	dir := "/static"
	handler := staticFileHandler(dir)

	if handler == nil {
		t.Error("expected handler to be non-nil")
	}

	// Simulate a request to the handler
	req, err := http.NewRequest("GET", dir+"/testfile.txt", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected status code 404, got %d", rr.Code)
	}
}
