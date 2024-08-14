package seqs

import (
	"slices"
	"testing"
)

func TestSkipUntil(t *testing.T) {
	var (
		ints    = Ints(1, 1)
		first10 = Limit(ints, 10)
		latter  = SkipUntil(first10, func(x int) bool { return x > 7 })
		got     = slices.Collect(latter)
		want    = []int{8, 9, 10}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
