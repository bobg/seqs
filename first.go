package seqs

import "iter"

// FirstUntil copies the input iterator to the output
// until the first element that causes f to return true.
// It is the same as FirstWhile(inp, notF),
// where notF is the boolean inverse of f.
func FirstUntil[T any](inp iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for val := range inp {
			if f(val) {
				return
			}
			if !yield(val) {
				return
			}
		}
	}
}

// FirstWhile copies the input iterator to the output
// while f returns true for the elements, then stops.
// It is the same as FirstUntil(inp, notF),
// where notF is the boolean inverse of f.
func FirstWhile[T any](inp iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return FirstUntil(inp, func(t T) bool { return !f(t) })
}

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
