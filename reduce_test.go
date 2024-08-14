package seqs

import (
	"fmt"
	"slices"
	"testing"
)

func TestReduce(t *testing.T) {
	cases := []struct {
		inp  []int
		want int
	}{{}, {
		inp:  []int{1},
		want: 1,
	}, {
		inp:  []int{1, 2},
		want: 3,
	}, {
		inp:  []int{1, 2, 3},
		want: 6,
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			var (
				inp = slices.Values(tc.inp)
				got = Reduce(inp, 0, func(a, b int) int { return a + b })
			)
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestReduce2(t *testing.T) {
	cases := []struct {
		inp  []int
		want int
	}{{}, {
		inp:  []int{1},
		want: 0,
	}, {
		inp:  []int{1, 2},
		want: 2,
	}, {
		inp:  []int{1, 2, 3},
		want: 8,
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			var (
				inp = slices.All(tc.inp)
				got = Reduce2(inp, 0, func(a, idx, val int) int { return a + (idx * val) })
			)
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
