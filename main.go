package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Feliz Final de semana!!")
}

func main() {
	http.HandleFunc("/", hello)

	http.ListenAndServe(":8080", nil)
}
