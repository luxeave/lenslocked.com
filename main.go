package main

import (
	"fmt"
	"net/http"
)

func someHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my bad site</h1>")
}

func main() {
	http.HandleFunc("/", someHandlerFunc)
	http.ListenAndServe(":3000", nil) // :3000 or localhost:3000 the same
}
