package main

import "fmt"

func main() {
	a := atomicBool{}
	fmt.Println(a.value())
	a.setValue(true)
	fmt.Println(a.value())

	s := safeBool{}
	fmt.Println(s.value())
	s.setValue(true)
	fmt.Println(s.value())
}
