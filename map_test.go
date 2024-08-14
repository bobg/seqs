package seqs

import (
	"slices"
	"testing"
)

func TestMap(t *testing.T) {
	var (
		inp  = slices.Values([]int{1, 2, 3, 4})
		m    = Map(inp, func(x int) int { return x * x })
		got  = slices.Collect(m)
		want = []int{1, 4, 9, 16}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMap2(t *testing.T) {
	var (
		inp   = slices.All([]int{1, 2, 3, 4})
		m     = Map2(inp, func(idx, x int) (int, int) { return idx, idx * x })
		pairs = ToPairs(m)
		got   = slices.Collect(pairs)
		want  = []Pair[int, int]{{0, 0}, {1, 2}, {2, 6}, {3, 12}}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
