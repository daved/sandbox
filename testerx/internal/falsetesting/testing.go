package falsetesting

import "fmt"

type T struct {
	Name string
}

func (t *T) Skip() {
	fmt.Printf("skipped %s\n", t.Name)
}
