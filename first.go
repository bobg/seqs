package seqs

import "iter"

// First returns the first value of seq and true.
// If seq is empty, it returns the zero value of T and false.
//
// Deprecated: This function can leak resources.
// Use [iter.Pull] instead.
func First[T any](seq iter.Seq[T]) (T, bool) {
	for v := range seq {
		return v, true
	}

	var zero T
	return zero, false
}

// First2 returns the first pair of values in seq, and true.
// If seq is empty, it returns the zero values of T and U, and false.
//
// Deprecated: This function can leak resources.
// Use [iter.Pull2] instead.
func First2[T, U any](seq iter.Seq2[T, U]) (T, U, bool) {
	for x, y := range seq {
		return x, y, true
	}

	var (
		zeroT T
		zeroU U
	)
	return zeroT, zeroU, false
}
