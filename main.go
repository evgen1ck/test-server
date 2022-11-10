package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет!")
}
