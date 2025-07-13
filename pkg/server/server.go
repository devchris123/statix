package server

import (
	"net/http"
)

// Server represents a simple HTTP server.
type server struct {
	addr   string
	server http.Server
}

// NewServer creates a new server instance with the specified address.
func NewServer(addr string) *server {
	return &server{addr: addr}
}

// Start starts the HTTP server.
func (s *server) Start() error {
	s.server.Addr = s.addr
	return s.server.ListenAndServe()
}
