package main

import "fmt"

func multicomp(s string) byte {
	if s[0] == 'n' || s[0] == 'N' {
		return 1
	}

	if s[0] == 'm' || s[0] == 'M' {
		return 1
	}

	return 0
}

func singlecomp(s string) byte {
	if s[0]|0x20 == 'n' {
		return 1
	}

	if s[0]|0x20 == 'm' {
		return 1
	}

	return 0
}

func main() {
	fmt.Println(multicomp("N"))
	fmt.Println(singlecomp("N"))
	fmt.Println(multicomp("M"))
	fmt.Println(singlecomp("M"))
	fmt.Println(multicomp("X"))
	fmt.Println(singlecomp("X"))
}
