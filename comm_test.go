package seqs

import (
	"slices"
	"testing"
)

func TestCommLeft(t *testing.T) {
	var (
		left  = slices.Values([]int{1, 2, 3, 4})
		right = slices.Values([]int{2, 4, 6, 8})
		got   = slices.Collect(CommLeft(left, right))
		want  = []int{1, 3}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCommRight(t *testing.T) {
	var (
		left  = slices.Values([]int{1, 2, 3, 4})
		right = slices.Values([]int{2, 4, 6, 8})
		got   = slices.Collect(CommRight(left, right))
		want  = []int{6, 8}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCommBoth(t *testing.T) {
	var (
		left  = slices.Values([]int{1, 2, 3, 4})
		right = slices.Values([]int{2, 4, 6, 8})
		got   = slices.Collect(CommBoth(left, right))
		want  = []int{2, 4}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
