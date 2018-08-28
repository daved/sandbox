package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go func() {
		time.Sleep(time.Millisecond * 5)
		fmt.Println("Calculated result!")
		c <- "42"
	}()

	fmt.Println("Waiting...")
	fmt.Println("The answer is: " + <-c)
}
