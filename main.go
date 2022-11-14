package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	http.HandleFunc("/greet/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/greet/"):]
		fmt.Fprintf(w, "Hello %s\n", name)
	})

	//http.ListenAndServe(":9990", nil
	http.ListenAndServeTLS(":9990",
		"/etc/letsencrypt/live/digitalshop.evgenick.com/fullchain.pem",
		"/etc/letsencrypt/live/digitalshop.evgenick.com/privkey.pem",
		nil)
}
