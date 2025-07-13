// Package server provides a simple HTTP server to serve static files.
package server

import (
	"log"
	"net/http"
	"path/filepath"
)

const defaultStaticDir = "./static"

type server struct {
	server http.Server
}

// NewServer creates a new server instance with the specified address.
func NewServer(addr string) *server {
	handler := http.DefaultServeMux

	handler.HandleFunc("/", staticFileHandler(defaultStaticDir))

	return &server{server: http.Server{
		Addr:    addr,
		Handler: handler,
	}}
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
