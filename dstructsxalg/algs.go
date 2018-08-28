package main

import "math"

// greatest common denominator
func gcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// least common multiple
func lcm(a, b uint64) uint64 {
	return (a / gcd(a, b)) * b
}

// factors (of number)
func fon(n uint64) []int {
	res := []int{}
	has := false

	for n%uint64(2) == 0 {
		has = true
		n = n / 2
	}

	if has {
		has = false
		res = append(res, 2)
	}

	f := uint64(3)
	stop := n * n
	for f <= stop {
		for n%f == 0 {
			has = true
			n = n / f
			stop = n * n
		}

		if has {
			res = append(res, int(f))
		}

		f = f + 2
	}

	if n > 1 {
		res = append(res, int(n))
	}

	return res
}

// sieve of eratosthenes
func soEratos(n uint64) []int {
	isPrime := make([]bool, n+1)

	isPrime[2] = true

	for i := uint64(3); i <= n; i = i + 2 {
		isPrime[i] = true
	}

	sqrt := func(n uint64) uint64 {
		f := float64(n)
		sr := math.Sqrt(f)
		sr += .5
		return uint64(sr)
	}

	for i := uint64(3); i <= sqrt(n); i = i + 2 {
		if isPrime[i] {
			for m := i * i; m <= n; m = m + 2*i {
				isPrime[m] = false
			}
		}
	}

	// process results
	res := []int{}
	for k, v := range isPrime {
		if v {
			res = append(res, k)
		}
	}
	return res
}
