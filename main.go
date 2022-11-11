package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func setDefaultHeaders(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloooooooo")
}

func main() {
	fmt.Println("helloooooo")
	r := mux.NewRouter()

	r.HandleFunc("/", setDefaultHeaders)
	http.ListenAndServe(":8000", r)
}
