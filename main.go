package main

import (
	"fmt"
	"github.com/leonardaustin/go-olympus/handler"
)

func main() {

	// Handle endpoint registration
	testvar := handler.EndpointList()
	fmt.Println(testvar)

	// Handel all http requests
	handler.Http()

	// Handle all proto requests
}

func test() {

}
