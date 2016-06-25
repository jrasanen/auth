package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/jrasanen/httpauth/auth"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "I like turtles!\n")
}

func main() {
	fmt.Println("Starting server")
	helloHandler := http.HandlerFunc(hello)
	http.Handle("/", auth.AuthorizeBasic(helloHandler))
	http.ListenAndServe(":8000", nil)
}
