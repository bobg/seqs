package seqs

import (
	"fmt"
	"slices"
	"testing"
)

func TestUniq(t *testing.T) {
	cases := []struct {
		inp  []int
		want []int
	}{{}, {
		inp:  []int{1},
		want: []int{1},
	}, {
		inp:  []int{1, 2},
		want: []int{1, 2},
	}, {
		inp:  []int{1, 2, 2},
		want: []int{1, 2},
	}, {
		inp:  []int{1, 2, 1},
		want: []int{1, 2, 1},
	}, {
		inp:  []int{1, 2, 1, 2},
		want: []int{1, 2, 1, 2},
	}, {
		inp:  []int{1, 2, 2, 1},
		want: []int{1, 2, 1},
	}, {
		inp:  []int{1, 2, 2, 1, 2},
		want: []int{1, 2, 1, 2},
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			var (
				inp = slices.Values(tc.inp)
				seq = Uniq(inp)
				got = slices.Collect(seq)
			)
			if !slices.Equal(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
