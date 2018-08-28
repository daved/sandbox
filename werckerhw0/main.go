package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codemodus/mixmux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World.")
}

func main() {
	m := mixmux.NewRouter(nil)

	m.Get("/api/hello", http.HandlerFunc(helloHandler))

	fmt.Fprintln(os.Stderr, http.ListenAndServe(":5454", m))
}
