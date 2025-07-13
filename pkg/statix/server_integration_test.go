package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// TestStaticFileHandler tests the staticFileHandler function.
// The test creates a dummy directory with dummy files and checks if the handler serves them correctly.
func TestIntegrationStaticFileHandler(t *testing.T) {
	dummyFile0 := "testfile0.txt"
	dummyFile1 := "testfile1.txt"
	dummyFile0Content := "This is a test file."
	dummyFile1Content := "This is another test file."
	readPerm := os.FileMode(0644)

	// Create a temporary directory for testing
	staticDir, err := os.MkdirTemp(".", "static_")
	if err != nil {
		t.Fatalf("could not create temp dir: %v", err)
	}
	t.Cleanup(func() {
		// Remove the temporary directory after the test
		if err := os.RemoveAll(staticDir); err != nil {
			t.Fatalf("could not remove temp dir: %v", err)
		}
	})
	// Create a dummy file in the temporary directory
	if err := os.WriteFile(filepath.Join(staticDir, dummyFile0), []byte(dummyFile0Content), readPerm); err != nil {
		t.Fatalf("could not create dummy file: %v", err)
	}
	// Create another dummy file
	if err := os.WriteFile(filepath.Join(staticDir, dummyFile1), []byte(dummyFile1Content), readPerm); err != nil {
		t.Fatalf("could not create another dummy file: %v", err)
	}

	handler := staticFileHandler(staticDir)

	if handler == nil {
		t.Error("expected handler to be non-nil")
	}

	// Simulate a request to the handler
	req, err := http.NewRequest("GET", dummyFile0, nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("expected status OK, got %d", rr.Code)
	}
	if rr.Body.String() != dummyFile0Content {
		t.Errorf("expected body %s, got %s", dummyFile0Content, rr.Body.String())
	}
}
