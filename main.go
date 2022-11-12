// package main
//
// import (
//
//	"fmt"
//	"net/http"
//
// )
//
//	func main() {
//		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//			fmt.Fprintf(w, "Hello World")
//		})
//
//		http.HandleFunc("/greet/", func(w http.ResponseWriter, r *http.Request) {
//			name := r.URL.Path[len("/greet/"):]
//			fmt.Fprintf(w, "Hello %s\n", name)
//		})
//
//		http.ListenAndServeTLS(":9990", "/etc/nginx/ssl/evgenick.pem", "/etc/nginx/ssl/evgenick.key", nil)
//	}
package main

import (
	"log"
	"net/http"
)

func HelloSSLServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func main() {
	http.HandleFunc("/", HelloSSLServer)
	go http.ListenAndServe("85.193.95.12:80", nil)
	err := http.ListenAndServeTLS("localhost:9990", "/etc/ssl/digitalshop-evgenick-com.crt", "/etc/ssl/digitalshop-evgenick-com.key", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
