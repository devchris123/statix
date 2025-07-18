package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestNewServer tests the NewServer function.
func TestNewServer(t *testing.T) {
	addr := "localhost:8080"
	srv, err := NewServer(
		WithPort(8080),
		WithStaticDir("./static"),
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

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

func TestWithHost(t *testing.T) {
	opts := &options{}
	host := "localhost"

	err := WithHost(host)(opts)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if opts.host == nil || *opts.host != host {
		t.Errorf("expected host %s, got %v", host, opts.host)
	}

	// Test empty host
	err = WithHost("")(opts)
	if err == nil || err.Error() != "host cannot be empty" {
		t.Errorf("expected error for empty host, got %v", err)
	}
}

func TestWithPort(t *testing.T) {
	opts := &options{}
	port := 8080

	err := WithPort(port)(opts)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if opts.port == nil || *opts.port != port {
		t.Errorf("expected port %d, got %v", port, opts.port)
	}

	// Test invalid ports
	err = WithPort(0)(opts)
	if err == nil || err.Error() != "invalid port number" {
		t.Errorf("expected error for invalid port, got %v", err)
	}

	err = WithPort(70000)(opts)
	if err == nil || err.Error() != "invalid port number" {
		t.Errorf("expected error for invalid port, got %v", err)
	}
}

func TestWithStaticDir(t *testing.T) {
	opts := &options{}
	dir := "/static"

	err := WithStaticDir(dir)(opts)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if opts.staticDir == nil || *opts.staticDir != dir {
		t.Errorf("expected static directory %s, got %v", dir, opts.staticDir)
	}

	// Test empty directory
	err = WithStaticDir("")(opts)
	if err == nil || err.Error() != "static directory cannot be empty" {
		t.Errorf("expected error for empty directory, got %v", err)
	}
}
