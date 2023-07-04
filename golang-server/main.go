package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter = 0

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request from:", r.RemoteAddr)
	response := fmt.Sprintf("You are number %d", counter)
	counter++
	fmt.Fprintf(w, response)
}

func main() {
	http.HandleFunc("/", helloHandler)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}