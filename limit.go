package seqs

import "iter"

// Limit returns an iterator over seq that stops after n values.
func Limit[V any](seq iter.Seq[V], n int) iter.Seq[V] {
	return func(yield func(V) bool) {
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

// Limit2 returns an iterator over seq that stops after n key-value pairs.
func Limit2[K, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
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
