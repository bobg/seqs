package seqs

import (
	"slices"
	"testing"
)

func TestResumable(t *testing.T) {
	var (
		inp  = slices.Values([]int{2, 3, 5, 7, 11, 13})
		stop func()
	)

	inp, stop = Resumable(inp)
	defer stop()

	var got1 []int
	for val := range inp {
		got1 = append(got1, val)
		if val == 5 {
			break
		}
	}

	var got2 []int
	for val := range inp {
		got2 = append(got2, val)
	}

	var (
		want1 = []int{2, 3, 5}
		want2 = []int{7, 11, 13}
	)

	if !slices.Equal(got1, want1) {
		t.Errorf("got %v, want %v", got1, want1)
	}
	if !slices.Equal(got2, want2) {
		t.Errorf("got %v, want %v", got2, want2)
	}
}
