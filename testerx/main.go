package main

import (
	"fmt"

	"github.com/daved/sandbox/testerx/internal/falsetesting"
	. "github.com/daved/sandbox/testerx/internal/ttype"
)

func main() {
	t := &falsetesting.T{}
	one(t)
	two(t)
	three(t)
}

func one(t *falsetesting.T) {
	t.Name = "one"

	RunForTypes(t, Critical)
	fmt.Println("one - not skipped for critical")
}

func two(t *falsetesting.T) {
	t.Name = "two"

	RunForTypes(t, Normal)
	fmt.Println("two - not skipped for normal")
}

func three(t *falsetesting.T) {
	fmt.Println("three - not skipped ever")
}

func four(t *falsetesting.T) {
	t.Name = "four"

	RunForTypes(t, Normal, Critical)
	fmt.Println("two - not skipped for normal or critical")
}
