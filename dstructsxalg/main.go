package main

import (
	"fmt"
	"time"
)

func main() {
	nowNano := uint64(time.Now().UnixNano())

	lcg := newLCG(nowNano)
	mcg := newMCG(nowNano)
	fmt.Printf("lcg rand = %d, mcg rand = %d\n", lcg.gen(), mcg.gen())
	//fmt.Printf("lcg runs to loop = %d, mcg runs to loop = %d\n", runsToLoop(lcg), runsToLoop(mcg))

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	fmt.Println("starting array:", a)
	randomizeArray(lcg, a, 10)
	fmt.Println("----randomized:", a)

	n0, n1 := lcg.gen(), mcg.gen()
	fmt.Printf("gcd of %d and %d = %d\n", n0, n1, gcd(n0, n1))
	n0, n1 = 32888301, 8567288682
	fmt.Printf("gcd of %d and %d = %d\n", n0, n1, gcd(n0, n1))

	n0, n1 = lcg.gen(), mcg.gen()
	fmt.Printf("lcm of %d and %d = %d\n", n0, n1, lcm(n0, n1))
	n0, n1 = 12, 15
	fmt.Printf("lcm of %d and %d = %d\n", n0, n1, lcm(n0, n1))

	n0 = 7524
	fmt.Printf("factors of %d = %v\n", n0, fon(n0))

	fmt.Printf("sieve of eratosthenes for %d = %v\n", n0, soEratos(n0))
}
