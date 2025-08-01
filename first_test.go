package seqs

import (
	"fmt"
	"slices"
	"testing"
)

func TestFirst(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		_, ok, stop := First(Empty[int])
		if ok {
			t.Error("got ok=true, want false")
		}
		stop()
	})

	t.Run("non-empty", func(t *testing.T) {
		ints := Ints(17, 1)
		got, ok, stop := First(ints)
		if !ok {
			t.Error("got ok=false, want true")
		} else if got != 17 {
			t.Errorf("got %v, want 17", got)
		}
		stop()
	})
}

func TestFirstUntil(t *testing.T) {
	cases := []struct {
		inp, want []int
	}{{
		inp:  nil,
		want: nil,
	}, {
		inp:  []int{4},
		want: []int{4},
	}, {
		inp:  []int{4, 5, 6, 7},
		want: []int{4, 5, 6, 7},
	}, {
		inp:  []int{4, 5, 6, 7, 8},
		want: []int{4, 5, 6, 7},
	}, {
		inp:  []int{8, 9, 10},
		want: nil,
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			gotSeq := FirstUntil(slices.Values(tc.inp), func(x int) bool { return x > 7 })
			gotSlice := slices.Collect(gotSeq)
			if !slices.Equal(gotSlice, tc.want) {
				t.Errorf("got %v, want %v", gotSlice, tc.want)
			}
		})
	}
}
