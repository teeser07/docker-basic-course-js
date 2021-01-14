// https://medium.com/faun/docker-multi-stage-build-for-go-a6821eabde1a

package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	serverAddress = ":3030"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("welcome")
	fmt.Fprintf(w, "Hello welcome %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Listening on port ", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, nil))
}
