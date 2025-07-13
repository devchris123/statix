package main

import (
	statix "github.com/devchris123/statix/pkg/statix"
)

func main() {
	// Create a new server instance with the specified address
	srv := statix.NewServer(":8080")

	// Start the server
	if err := srv.Start(); err != nil {
		panic(err)
	}
}
