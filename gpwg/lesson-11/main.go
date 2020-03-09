package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var (
		dtxt = "CSOITEUIWUIZNSROCNKFD"
		dkey = "GOLANG"
		ctxt = "MYMESSAGEGOESHEREANDITISEXTRALONGANDMOREWITHUSELESSSTUFF"
		ckey = "GETPROGRAMMINGWITHGO"
	)

	fmt.Println(decipher(dkey, dtxt))

	c := cipher(ckey, ctxt)
	fmt.Println(c)
	fmt.Println(decipher(ckey, c))
}

func cipher(key, txt string) string {
	rs := make([]rune, utf8.RuneCountInString(txt))
	off := 'A'
	cnt := 'Z' + 1 - off

	for i, r := range txt {
		k, _ := utf8.DecodeRuneInString(key[i%len(key):])
		k -= off
		r += k
		r -= off
		r %= cnt
		r += off
		rs[i] = r
	}

	return string(rs)
}

func decipher(key, txt string) string {
	rs := make([]rune, utf8.RuneCountInString(txt))
	off := 'A'
	cnt := 'Z' + 1 - off

	for i, r := range txt {
		k, _ := utf8.DecodeRuneInString(key[i%len(key):])
		r -= k
		r += cnt
		r %= cnt
		r += off
		rs[i] = r
	}

	return string(rs)
}
