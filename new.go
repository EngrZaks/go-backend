package main

import (
	"fmt"
	"net/http"
)

func main() {
	// fmt.Println("hello")
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "hellow world")
	})
	http.ListenAndServe(":8000", nil)
}
