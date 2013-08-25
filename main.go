package main

import (
	"fmt"
	"github.com/leonardaustin/go-olympus/endpoint"
	"github.com/leonardaustin/go-olympus/server"
)

func main() {

	// Handle endpoint registration
	testvar := endpoint.List()
	fmt.Println(testvar)

	// Handel all http requests
	server.Http()

	// Handle all internal proto requests
}

func test() {

}
