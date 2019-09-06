package main

import (
	"fmt"
	"net"
	"net/http"
)

// StartServer begins the listener
func StartServer() {
	SetupHandlers()
	listenAddr := net.JoinHostPort("0.0.0.0", "8080")
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		fmt.Printf("Error from ListenAndServer %v\n", err)
	}
}

// SetupHandlers adds the handlers in.
func SetupHandlers() {
	http.HandleFunc("/show/", ShowParams)
}

// ShowParams displays the url parameters
func ShowParams(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got url=%v\n", r.URL)
}

func main() {
	StartServer()
}
