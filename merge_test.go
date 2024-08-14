package seqs

import (
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	var (
		slice1 = []int{3, 6, 9}
		slice2 = []int{1, 4, 7, 10}
		slice3 = []int{2, 5, 8}
	)

	t.Run("none", func(t *testing.T) {
		var (
			m   = MergeAll[int]()
			got = slices.Collect(m)
		)
		if len(got) != 0 {
			t.Errorf("got %v, want []", got)
		}
	})

	t.Run("one", func(t *testing.T) {
		var (
			seq1 = slices.Values(slice1)
			m    = MergeAll(seq1)
			got  = slices.Collect(m)
		)
		if !slices.Equal(got, slice1) {
			t.Errorf("got %v, want %v", got, slice1)
		}
	})

	t.Run("two", func(t *testing.T) {
		var (
			seq1 = slices.Values(slice1)
			seq2 = slices.Values(slice2)
			m    = MergeAll(seq1, seq2)
			got  = slices.Collect(m)
			want = []int{1, 3, 4, 6, 7, 9, 10}
		)
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("three", func(t *testing.T) {
		var (
			seq1 = slices.Values(slice1)
			seq2 = slices.Values(slice2)
			seq3 = slices.Values(slice3)
			m    = MergeAll(seq1, seq2, seq3)
			got  = slices.Collect(m)
			want = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		)
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestMerge2(t *testing.T) {
	var (
		odds  = slices.Values([]Pair[int, string]{{1, "one"}, {3, "three"}, {5, "five"}})
		evens = slices.Values([]Pair[int, string]{{2, "two"}, {4, "four"}, {6, "six"}})
		m     = Merge2(FromPairs(odds), FromPairs(evens))
		names = Right(m)
		got   = slices.Collect(names)
		want  = []string{"one", "two", "three", "four", "five", "six"}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
