package main

import (
	"fmt"
	"net/http"
	"transliterate"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, What are you searching for?",
		"1. /transliterate",
		"2. /questions")
	fmt.Println(r)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/transliterate", transliterate.HandleTransliterateQuery)
	http.ListenAndServe(":8080", nil)
}
