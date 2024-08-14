package seqs

import "iter"

// Filter returns an iterator over seq that includes only the values v for which f(v) is true.
func Filter[V any, F ~func(V) bool](seq iter.Seq[V], f F) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if f(v) && !yield(v) {
				return
			}
		}
	}
}

// Filter2 returns an iterator over seq that includes only the pairs t,u for which f(t, u) is true.
func Filter2[T, U any, F ~func(T, U) bool](seq iter.Seq2[T, U], f F) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		for t, u := range seq {
			if f(t, u) && !yield(t, u) {
				return
			}
		}
	}
}
