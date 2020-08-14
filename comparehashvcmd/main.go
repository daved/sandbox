package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"log"
	"os/exec"
	"time"

	"golang.org/x/crypto/blake2b"
)

func main() {
	h, err := blake2b.New256(nil)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer([]byte("test this out"))
	buf2 := bytes.NewBuffer([]byte("test this out"))

	start := time.Now()
	out := dataHash2(h, buf.Bytes())
	dur := time.Since(start).Microseconds()
	fmt.Println(dur)
	fmt.Println(out)

	buf.Reset()

	start = time.Now()
	out = dataHash2(h, buf2.Bytes())
	dur = time.Since(start).Microseconds()
	fmt.Println(dur)
	fmt.Println(out)

	start = time.Now()
	out = callCmd()
	dur = time.Since(start).Microseconds()
	fmt.Println(dur)
	fmt.Println(out)

	start = time.Now()
	out = callCmd()
	dur = time.Since(start).Microseconds()
	fmt.Println(dur)
	fmt.Println(out)
}

func dataHash(h hash.Hash, buf io.Reader) string {
	h.Reset()
	if _, err := io.Copy(h, buf); err != nil {
		log.Fatal(err)
	}

	hash := h.Sum(nil)

	return hex.EncodeToString(hash)
}

func dataHash2(h hash.Hash, data []byte) string {
	h.Reset()
	hash := h.Sum(data)

	return hex.EncodeToString(hash)
}

func callCmd() string {
	cmd := exec.Command("ls", "-l")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	return out.String()
}
