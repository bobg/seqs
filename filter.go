package seqs

import "iter"

// Filter returns an iterator over seq that only includes
// the values v for which f(v) is true.
func Filter[V any, F ~func(V) bool](seq iter.Seq[V], f F) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if f(v) && !yield(v) {
				return
			}
		}
	}
}

// Filter2 returns an iterator over seq that only includes
// the pairs k, v for which f(k, v) is true.
func Filter2[K, V any, F ~func(K, V) bool](seq iter.Seq2[K, V], f F) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if f(k, v) && !yield(k, v) {
				return
			}
		}
	}
}
