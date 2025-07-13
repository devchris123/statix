package main

import (
	statix "github.com/devchris123/statix/pkg/statix"
)

func main() {
	// Create a new server instance with the specified address
	srv, err := statix.NewServer(
		statix.WithPort(8080),
		statix.WithStaticDir("./html"),
	)

	if err != nil {
		panic(err) // Handle error appropriately
	}

	// Start the server
	if err := srv.Start(); err != nil {
		panic(err)
	}
}
