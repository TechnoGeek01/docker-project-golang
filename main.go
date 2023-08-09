package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func myFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

}

func myFunc2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")

}

func main() {
	http.HandleFunc("/", myFunc)
	http.HandleFunc("/hi", myFunc2)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
