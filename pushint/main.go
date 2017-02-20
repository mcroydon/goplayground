package main

import (
	"fmt"
	"io"
	"net/http"
	"log"
	"math/rand"
)

// This handler will push a pseudo-random integer via /int/ if possible.
func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("handling /\n")
	if p, ok := w.(http.Pusher); ok {
		log.Print("Pushing.\n")
		p.Push("/int/", nil)
	} else {
		log.Print("Push not available.\n")
	}
	io.WriteString(w, "Hello there.")
}

// Serve a pseudo-random integer.
func intHandler(w http.ResponseWriter, r *http.Request) {
	i := rand.Int()
	log.Printf("Handling /int/ with %v\n", i)
	fmt.Fprintf(w, "%v", i)
}

// Start a server on port 8443 with a self-signed certificate.
// To test, visit https://localhost:8443. If your browser supports HTTP/2 push,
// you will see the pushed value if you then visit https://localhost:8443/int/.
// If push is not supported, a new request will be made and you will see a new
// pseudo-random integer instead.
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/int/", intHandler)
	err := http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
