package seqs

import (
	"slices"
	"testing"
)

func TestCheckEmpty(t *testing.T) {
	var (
		ints   = Ints(1, 1)
		first3 = Limit(ints, 3)
	)
	first3, isEmpty := CheckEmpty(first3)
	if isEmpty {
		t.Error("got isEmpty == true, want false")
	}
	got := slices.Collect(first3)
	want := []int{1, 2, 3}
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	_, isEmpty = CheckEmpty(Empty[int])
	if !isEmpty {
		t.Error("got isEmpty == false, want true")
	}
}

func TestCheckEmpty2(t *testing.T) {
	var (
		ints    = Ints(1, 1)
		squares = Map(Ints(1, 1), func(i int) int { return i * i })
		zipped  = ZipVals(ints, squares)
		first3  = Limit2(zipped, 3)
	)
	first3, isEmpty := CheckEmpty2(first3)
	if isEmpty {
		t.Error("got isEmpty == true, want false")
	}
	got := slices.Collect(ToPairs(first3))
	want := []Pair[int, int]{{1, 1}, {2, 4}, {3, 9}}
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	_, isEmpty = CheckEmpty2(Empty2[int, int])
	if !isEmpty {
		t.Error("got isEmpty == false, want true")
	}
}
