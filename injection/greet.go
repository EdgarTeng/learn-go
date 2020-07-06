package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet means say Hello for a name
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreetHandler using http way to say hello
func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler))
}
