package seqs

import (
	"slices"
	"testing"
)

func TestConcat(t *testing.T) {
	var (
		a    = slices.Values([]int{1, 2, 3})
		b    = slices.Values([]int{4, 5, 6})
		c    = Concat(a, b)
		got  = slices.Collect(c)
		want = []int{1, 2, 3, 4, 5, 6}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestConcat2(t *testing.T) {
	var (
		a1      = slices.Values([]int{1, 2, 3})
		b1      = slices.Values([]int{4, 5, 6})
		zipped1 = Zip(a1, b1)
		a2      = slices.Values([]int{100, 101, 102})
		b2      = slices.Values([]int{103, 104, 105})
		zipped2 = Zip(a2, b2)
		c       = Concat2(zipped1, zipped2)
		pairs   = ToPairs(c)
		got     = slices.Collect(pairs)
		want    = []Pair[int, int]{
			{X: 1, Y: 4},
			{X: 2, Y: 5},
			{X: 3, Y: 6},
			{X: 100, Y: 103},
			{X: 101, Y: 104},
			{X: 102, Y: 105},
		}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
