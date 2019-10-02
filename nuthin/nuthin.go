// Package nuthin provides a trivial primitive and function for demonstrating
// table-driven testing, documentation, and godoc example.
//
// If needed, install godoc by running `go get golang.org/x/tools/cmd/godoc`.
// Then run `godoc` from inside this directory and open your browser to
// `http://localhost:6060/pkg/github.com/daved/sandbox/nuthin`.
package nuthin

// Up describes what's up.
type Up struct {
	N int
}

// Much returns the greater of two instances of Up.
func Much(a, b Up) Up {
	c := a
	if b.N > a.N {
		c = b
	}
	return c
}
