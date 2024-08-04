package seqs

import (
	"context"
	"errors"
	"iter"
	"slices"
	"testing"
)

func TestToChan(t *testing.T) {
	var (
		ints  = Ints(0, 1)
		first = FirstN(ints, 10)
		ch    = ToChan(first)
		got   []int
		want  = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	)
	for val := range ch {
		got = append(got, val)
	}
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestToChanContext(t *testing.T) {
	var (
		ch1  = make(chan int, 1)
		seq1 = FromChan(ch1)
		ctx  = context.Background()
	)
	ctx, cancel := context.WithCancel(ctx)
	ch2, errptr := ToChanContext(ctx, seq1)

	ch1 <- 1
	got, ok := <-ch2
	if !ok {
		t.Fatal("channel closed")
	}
	if got != 1 {
		t.Errorf("got %d, want 1", got)
	}

	ch1 <- 2
	got, ok = <-ch2
	if !ok {
		t.Fatal("channel closed")
	}
	if got != 2 {
		t.Errorf("got %d, want 2", got)
	}

	cancel()

	ch1 <- 3

	if _, ok := <-ch2; ok {
		t.Fatal("channel still open after context cancellation")
	}

	if !errors.Is(*errptr, context.Canceled) {
		t.Errorf("got error %v, want %v", *errptr, context.Canceled)
	}
}

func TestFromChan(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	var (
		seq  = FromChan(ch)
		got  = slices.Collect(seq)
		want = []int{0, 1, 2}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFromChanContext(t *testing.T) {
	t.Run("canceled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 1; i <= 10; i++ {
				select {
				case <-ctx.Done():
					return
				case ch <- i:
				}
			}
		}()

		seq, errptr := FromChanContext(ctx, ch)
		next, stop := iter.Pull(seq)
		defer stop()
		if _, ok := next(); !ok {
			t.Fatal("no first value in sequence")
		}

		cancel()

		if _, ok := next(); ok {
			t.Fatal("next value available after context cancellation")
		}

		stop()

		err := *errptr
		if !errors.Is(err, context.Canceled) {
			t.Errorf("got error %v, want %v", err, context.Canceled)
		}
	})

	t.Run("uncanceled", func(t *testing.T) {
		ch := make(chan int)
		go func() {
			for i := 1; i <= 10; i++ {
				ch <- i
			}
			close(ch)
		}()

		seq, errptr := FromChanContext(context.Background(), ch)
		got := slices.Collect(seq)
		if *errptr != nil {
			t.Fatal(*errptr)
		}
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestGo2(t *testing.T) {
	seq, errptr := Go2(func(ch chan<- Pair[int, int]) error {
		ch <- Pair[int, int]{X: 1, Y: 2}
		ch <- Pair[int, int]{X: 3, Y: 4}
		return nil
	})
	pairs := ToPairs(seq)
	if *errptr != nil {
		t.Fatal(*errptr)
	}
	got := slices.Collect(pairs)
	want := []Pair[int, int]{{X: 1, Y: 2}, {X: 3, Y: 4}}
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
