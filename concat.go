package seqs

import "iter"

// Concat concatenates the members of the input iterators.
func Concat[T any](inps ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, inp := range inps {
			if inp == nil {
				continue
			}
			for val := range inp {
				if !yield(val) {
					return
				}
			}
		}
	}
}

func Concat2[T, U any](inps ...iter.Seq2[T, U]) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		for _, inp := range inps {
			if inp == nil {
				continue
			}
			for x, y := range inp {
				if !yield(x, y) {
					return
				}
			}
		}
	}
}
