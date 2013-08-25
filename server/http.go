package server

import (
	"fmt"
	"log"
	"net/http"
)

// Handles all http requests
func Http() {
	http.HandleFunc("/", httpRequest)
	http.ListenAndServe(":8080", nil)

}

func httpRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	log.Println(r.URL)
}
