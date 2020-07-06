package main

import (
	"fmt"
	"net/http"
)

// PlayerServer is http server
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
