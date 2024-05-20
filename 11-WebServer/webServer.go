package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Apa kabar!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Halo!")
	})

	http.HandleFunc("/index", index)

	fmt.Println("Starting web server at http://localhost:9090/")

	http.ListenAndServe(":9090", nil)
}
