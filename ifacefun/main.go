package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	fn := func(s string) string {
		return strings.Replace(s, "e", "3", -1)
	}
	fnh := myFunc(fn)
	fnh.ServeHTTP(nil, nil)
	/// /// ///

	h := &hello{name: "sally"}

	m := http.NewServeMux()
	m.Handle("/hello", filterLs(h))
	m.Handle("/test", http.HandlerFunc(handleTest))

	fmt.Fprintln(os.Stderr, http.ListenAndServe(":1412", m))
}

type hello struct {
	name string
}

func (h *hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi, my name is %s\n", h.name)
}

type myFunc func(string) string

func (f myFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(os.Stdout, f("hello from myfunc"))
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello from handleTest")
}

func filterLs(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f := newFilterLetter(w, "l", "|")

		next.ServeHTTP(f, r)
	})
}

type filterLetter struct {
	http.ResponseWriter
	in  []byte
	out []byte
}

func newFilterLetter(w http.ResponseWriter, in, out string) *filterLetter {
	return &filterLetter{
		ResponseWriter: w,
		in:             []byte(in),
		out:            []byte(out),
	}
}

func (f *filterLetter) Write(d []byte) (int, error) {
	d = bytes.ReplaceAll(d, f.in, f.out)
	return f.ResponseWriter.Write(d)
}
