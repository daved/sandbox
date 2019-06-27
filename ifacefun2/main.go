package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/daved/sandbox/ifacefun2/mauth"
)

func main() {
	var (
		authKey = "auth"
		ma      = mauth.New("fred", "george", "sally")
	)

	m := http.NewServeMux()
	m.HandleFunc("/example", handleExample)

	authorize := authorizeFunc(authKey, ma)

	var h http.Handler
	h = http.HandlerFunc(handleExample)
	h = authorize(h)

	m.Handle("/wrap", authorize(http.HandlerFunc(handleExample)))

	if err := http.ListenAndServe(":1412", m); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
