package seqs

import "iter"

// Limit returns an iterator over the elements of seq that stops after n values
// (or when seq is exhausted, whichever comes first).
func Limit[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		if n <= 0 {
			return
		}
		for v := range seq {
			if !yield(v) {
				return
			}
			if n--; n <= 0 {
				break
			}
		}
	}
}

// Limit returns an iterator over the elements of seq that stops after n pairs
// (or when seq is exhausted, whichever comes first).
func Limit2[T, U any](seq iter.Seq2[T, U], n int) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		if n <= 0 {
			return
		}
		for k, v := range seq {
			if !yield(k, v) {
				return
			}
			if n--; n <= 0 {
				break
			}
		}
	}
}
