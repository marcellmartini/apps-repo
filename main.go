package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
	fmt.Printf("Bingo")
}

func main() {
	http.HandleFunc("/", hello)

	http.ListenAndServe(":8080", nil)
}
