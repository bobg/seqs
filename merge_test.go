package seqs

import (
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	var (
		mod0   = slices.Values([]int{3, 6, 9})
		mod1   = slices.Values([]int{1, 4, 7, 10})
		mod2   = slices.Values([]int{2, 5, 8})
		merged = MergeAll(mod0, mod1, mod2, Empty[int])
		got    = slices.Collect(merged)
		want   = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
