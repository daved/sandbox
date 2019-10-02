package nuthin_test

import (
	"testing"

	"github.com/daved/sandbox/nuthin"
)

func TestMuch(t *testing.T) {
	tests := map[string]struct {
		a, b int
		want int
	}{
		"pos a>":     {10, 9, 10},
		"neg a>":     {-9, -10, -9},
		"pos a<":     {2, 3, 3},
		"neg a<":     {-3, -2, -2},
		"pos a=":     {3, 3, 3},
		"neg a=":     {-3, -3, -3},
		"pos neg a>": {6, -6, 6},
		"neg pos a<": {-7, 7, 7},
		"zero a a>":  {0, -1, 0},
		"zero a a<":  {0, 1, 1},
		"zero b b>":  {-1, 0, 0},
		"zero b b<":  {1, 0, 1},
	}

	for tn, tt := range tests {
		want := nuthin.Up{tt.want}
		a, b := nuthin.Up{tt.a}, nuthin.Up{tt.b}
		got := nuthin.Much(a, b)
		if got != want {
			t.Errorf("%s: got %v, want %v", tn, got, want)
		}
	}
}
