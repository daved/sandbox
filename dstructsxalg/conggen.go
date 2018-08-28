package main

func init() {
	// silence
	_ = runsToLoop
}

type generator interface {
	gen() uint64
}

func runsToLoop(g generator) uint64 {
	start := g.gen()

	for i := uint64(1); ; i++ {
		if start == g.gen() {
			return i
		}
	}
}

func randomizeArray(g generator, a []int, n int) {
	if n > len(a) {
		n = len(a)
	}

	for k := range a[:n] {
		i := g.gen() % uint64(n)
		a[k], a[i] = a[i], a[k]
	}
}

// linear congruential generator
type lcg struct {
	val uint64
}

func newLCG(seed uint64) *lcg {
	return &lcg{val: seed}
}

func (g *lcg) gen() uint64 {
	g.val = (g.val*214013 + 2531011) % 4294967296
	return g.val
}

// multiplicative congruential generator
type mcg struct {
	val uint64
}

func newMCG(seed uint64) *mcg {
	return &mcg{val: seed}
}

func (g *mcg) gen() uint64 {
	g.val = (g.val * 214013) % 4294967296
	return g.val
}
