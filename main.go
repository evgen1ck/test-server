package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Internet! Please follow me.\n")
	})

	http.ListenAndServe(":80", nil)
}
