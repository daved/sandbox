package main

import (
	"fmt"
	"os"

	"github.com/daved/encdecfun/aesencdec"
)

func exit(err error) {
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	var err error
	defer func() { exit(err) }()

	key := []byte("testthistestthat")
	aesed, err := aesencdec.New(key)
	if err != nil {
		fmt.Printf("cannot get new aes encdec: %s", err)
		return
	}

	msg := "This is a message."
	fmt.Println(msg)

	enc, err := aesed.Encrypt([]byte(msg))
	if err != nil {
		fmt.Printf("cannot encrypt: %s\n", err)
		return
	}
	fmt.Println(string(enc))

	dec, err := aesed.Decrypt(enc)
	if err != nil {
		fmt.Printf("cannot decrypt: %s\n", err)
		return
	}
	fmt.Println(string(dec))

	if len(os.Args) > 1 && len(os.Args[1]) > 0 {
		out, err := aesed.Decrypt([]byte(os.Args[1]))
		if err != nil {
			fmt.Printf("cannot decrypt user input: %s", err)
			return
		}
		fmt.Println(string(out))
	}
}
