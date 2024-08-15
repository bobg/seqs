package seqs

import "testing"

func TestFirst(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		if _, ok := First(Empty[int]); ok {
			t.Error("got ok=true, want false")
		}
	})

	t.Run("non-empty", func(t *testing.T) {
		ints := Ints(17, 1)
		got, ok := First(ints)
		if !ok {
			t.Error("got ok=false, want true")
		} else if got != 17 {
			t.Errorf("got %v, want 17", got)
		}
	})
}
