package seqs

import (
	"slices"
	"testing"
)

func TestPeek(t *testing.T) {
	var (
		ints   = Ints(1, 1)
		first3 = Limit(ints, 3)
	)
	got, ok, orig := Peek(first3)
	if !ok {
		t.Error("got !ok, want ok")
	}
	if got != 1 {
		t.Errorf("got %d, want 1", got)
	}
	gotOrig := slices.Collect(orig)
	wantOrig := []int{1, 2, 3}
	if !slices.Equal(gotOrig, wantOrig) {
		t.Errorf("got %v, want %v", gotOrig, wantOrig)
	}

	_, ok, _ = Peek(Empty[int])
	if ok {
		t.Error("got ok, want !ok")
	}
}

func TestPeek2(t *testing.T) {
	var (
		ints    = Ints(1, 1)
		squares = Map(Ints(1, 1), func(i int) int { return i * i })
		zipped  = ZipVals(ints, squares)
		first3  = Limit2(zipped, 3)
	)
	gotA, gotB, ok, orig := Peek2(first3)
	if !ok {
		t.Error("got !ok, want ok")
	}
	if gotA != 1 || gotB != 1 {
		t.Errorf("got %d,%d, want 1,1", gotA, gotB)
	}
	gotOrig := slices.Collect(ToPairs(orig))
	wantOrig := []Pair[int, int]{{1, 1}, {2, 4}, {3, 9}}
	if !slices.Equal(gotOrig, wantOrig) {
		t.Errorf("got %v, want %v", gotOrig, wantOrig)
	}

	_, _, ok, _ = Peek2(Empty2[int, int])
	if ok {
		t.Error("got ok, want !ok")
	}
}
