package main

import (
	"fmt"
	"io/ioutil"
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
	http.HandleFunc("/", ShowDetails)
}

// ShowDetails displays the url parameters
func ShowDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v %v %v\n", r.Proto, r.Method, r.URL)
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading Body: %v\n", err)
			return
		}
		fmt.Printf("%v\n\n", string(body))
	}
	fmt.Printf("---\n")

}

func main() {
	StartServer()
}
