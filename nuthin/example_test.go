package nuthin_test

import (
	"fmt"

	"github.com/daved/sandbox/nuthin"
)

func ExampleMuch() {
	a := nuthin.Up{13}
	b := nuthin.Up{42}

	whatsup := nuthin.Much(a, b)
	fmt.Println(whatsup.N)

	// Output: 42
}
