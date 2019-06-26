package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world</h1>")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server starting...")
	http.ListenAndServe(":2525", nil)
}
