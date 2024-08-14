package seqs

import (
	"slices"
	"testing"
)

func TestFilter(t *testing.T) {
	var (
		ints   = Ints(1, 1)
		evens  = Filter(ints, func(n int) bool { return n%2 == 0 })
		first3 = Limit(evens, 3)
		got    = slices.Collect(first3)
		want   = []int{2, 4, 6}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, %v", got, want)
	}
}

func TestFilter2(t *testing.T) {
	var (
		nums     = slices.Values([]int{1, 3, 2, 4, 5, 7, 6, 8})
		enum     = Enumerate(nums)
		filtered = Filter2(enum, func(idx, n int) bool { return (idx*n)%2 == 0 })
		pairs    = ToPairs(filtered)
		got      = slices.Collect(pairs)
		want     = []Pair[int, int]{{0, 1}, {2, 2}, {3, 4}, {4, 5}, {6, 6}, {7, 8}}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
