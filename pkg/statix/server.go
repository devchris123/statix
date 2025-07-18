// Package server provides a simple HTTP server to serve static files.
package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

const defaultStaticDir = "./static"

type server struct {
	server http.Server
}

type options struct {
	host      *string
	port      *int
	staticDir *string
}

type Option func(*options) error

// WithHost sets the host for the server.
func WithHost(host string) Option {
	return func(opts *options) error {
		if host == "" {
			return errors.New("host cannot be empty") // Empty host
		}
		opts.host = &host
		return nil
	}
}

// WithPort sets the port for the server.
func WithPort(port int) Option {
	return func(opts *options) error {
		if port < 1 || port > 65535 {
			return errors.New("invalid port number") // Invalid port number
		}
		opts.port = &port
		return nil
	}
}

// WithStaticDir sets the directory to serve static files from.
func WithStaticDir(dir string) Option {
	return func(opts *options) error {
		if dir == "" {
			return errors.New("static directory cannot be empty") // Empty directory path
		}
		opts.staticDir = &dir
		return nil
	}
}

// NewServer creates a new server instance with the specified address.
func NewServer(opts ...Option) (*server, error) {
	// Apply options
	var o options
	for _, opt := range opts {
		if err := opt(&o); err != nil {
			return nil, err // Return error if option fails
		}
	}

	// Apply default options if not set
	applyDefaultOptions(&o)

	handler := http.DefaultServeMux
	handler.HandleFunc("/", staticFileHandler(*o.staticDir))

	return &server{server: http.Server{
		Addr:    fmt.Sprintf("%s:%d", *o.host, *o.port),
		Handler: handler,
	}}, nil
}

func applyDefaultOptions(opts *options) {
	if opts.host == nil {
		defaultHost := "localhost"
		opts.host = &defaultHost // Default host if not set
	}
	if opts.port == nil {
		defaultPort := 8080
		opts.port = &defaultPort // Default port if not set
	}
	if opts.staticDir == nil {
		defaultStaticDir := defaultStaticDir
		opts.staticDir = &defaultStaticDir // Default static directory
	}
}

func staticFileHandler(dir string) http.HandlerFunc {
	log.Printf("statix: Serving static files from %s", dir)
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("statix: Received request for %s", r.URL.Path)
		http.ServeFile(w, r, filepath.Join(dir, r.URL.Path))
	}
}

// Start starts the HTTP server.
func (s *server) Start() error {
	log.Printf("statix: Starting server on %s", s.server.Addr)
	return s.server.ListenAndServe()
}
