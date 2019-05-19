package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About page</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Contact page</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	fmt.Println("Server running.....")
	http.ListenAndServe(":4040", nil)
}
