package testassertion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertSubtract(t *testing.T) {
	ds := []struct {
		name           string
		min, sub, diff int
	}{
		{"zeroes", 0, 0, 0},
		{"tens", 10, 10, 0},
		{"larger min", 10, 1, 9},
		{"larger sub", 1, 10, -9},
		{"fail", 10, 10, -999},
	}

	calc := NewCalculator(nil)

	for _, d := range ds {
		t.Run(d.name, func(t *testing.T) {
			assert.Equal(t, d.diff, calc.Subtract(d.min, d.sub))
		})
	}
}

func TestStandardSubtract(t *testing.T) {
	ds := []struct {
		name           string
		min, sub, want int
	}{
		{"zeroes", 0, 0, 0},
		{"tens", 10, 10, 0},
		{"larger min", 10, 1, 9},
		{"larger sub", 1, 10, -9},
		{"fail", 10, 10, -999},
	}

	calc := NewCalculator(nil)

	for _, d := range ds {
		t.Run(d.name, func(t *testing.T) {
			got := calc.Subtract(d.min, d.sub)
			if got != d.want {
				t.Errorf("got %v, want %v", got, d.want)
			}
		})
	}
}

func TestStandardAltsSubtract(t *testing.T) {
	gwf := "got %v, want %v" // as global

	got := NewCalculator(nil).Subtract(10, 5)
	want := 5
	if got != want {
		t.Errorf(gwf, got, want)
	}

	// OR

	errf := func(t *testing.T, got, want interface{}) { // as global
		t.Errorf("got %v, want %v", got, want)
	}

	got = NewCalculator(nil).Subtract(10, 5)
	want = 5
	if got != want {
		errf(t, got, want)
	}
}
