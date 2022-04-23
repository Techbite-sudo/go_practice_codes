package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Index page</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About page</h1>")
}
func source(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h2>This is the source link. click to open the page.</h2>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/source", source)
	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", nil)
}
