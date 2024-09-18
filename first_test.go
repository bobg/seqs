package seqs

import "testing"

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
