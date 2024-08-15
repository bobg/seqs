package seqs

import "iter"

// First returns the first value of seq and true.
// If seq is empty, it returns the zero value of T and false.
func First[T any](seq iter.Seq[T]) (T, bool) {
	for v := range seq {
		return v, true
	}

	var zero T
	return zero, false
}
