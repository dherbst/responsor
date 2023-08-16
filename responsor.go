package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

// StartServer begins the listener
func StartServer(port string) {
	SetupHandlers()
	listenAddr := net.JoinHostPort("0.0.0.0", port)
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

	// print headers
	for h, v := range r.Header {
		for _, val := range v {
			fmt.Printf("%v: %v\n", h, val)
		}
	}

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading Body: %v\n", err)
			return
		}
		fmt.Printf("%v\n", string(body))
	}
	fmt.Printf("---end-of-request---\n")

}

func main() {

	port := "8899"
	fmt.Printf("Starting port=%v\n", port)
	StartServer(port)
}
