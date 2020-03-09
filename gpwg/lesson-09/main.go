package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var (
		neg3Msg   = "L fdph, L vdz, L frqtxhuhg."
		cipherMsg = "Hola Estación Espacial Internacional"
	)

	fmt.Println(decipherNeg3(neg3Msg))

	fmt.Println(rot13(cipherMsg))
}

func rot13(msg string) string {
	rs := make([]rune, utf8.RuneCountInString(msg))

	for _, r := range msg {
		var max rune
		switch {
		case r >= 'a' && r <= 'z':
			max = 'z'
		case r >= 'A' && r <= 'Z':
			max = 'Z'
		default:
			rs = append(rs, r)
			continue
		}

		r += 13
		if r > max {
			r -= 26
		}

		rs = append(rs, r)
	}

	return string(rs)
}

func decipherNeg3(msg string) string {
	bs := make([]byte, len(msg))

	for i := 0; i < len(msg); i++ {
		b := msg[i]

		var min byte
		switch {
		case b >= 'a' && b <= 'z':
			min = 'a'
		case b >= 'A' && b <= 'Z':
			min = 'A'
		default:
			bs = append(bs, b)
			continue
		}

		b -= 3
		if b < min {
			b += 26
		}

		bs = append(bs, b)
	}

	return string(bs)
}
