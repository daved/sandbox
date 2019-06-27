package main

import (
	"fmt"
	"net/http"
)

func handleExample(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is my response\n")
}
