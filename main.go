package main

import (
	"github.com/leonardaustin/go-olympus/endpoint"
	"github.com/leonardaustin/go-olympus/server"
)

func main() {

	// Handle endpoint registration
	endpoint.Listen()

	// Handel all http requests
	server.Http()

	// Handle all internal proto requests
}
