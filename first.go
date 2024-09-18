package seqs

import "iter"

// First returns the first value of seq and true.
// If seq is empty, it returns the zero value of T and false.
func First[T any](seq iter.Seq[T]) (T, bool, func()) {
	var stop func()
	seq, stop = Resumable(seq)
	for v := range seq {
		return v, true, stop
	}

	var zero T
	return zero, false, func() {}
}

// First2 returns the first pair of values of seq, and true.
// If seq is empty, it returns the zero values of T and U, and false.
func First2[T, U any](seq iter.Seq2[T, U]) (T, U, bool, func()) {
	var stop func()
	seq, stop = Resumable2(seq)
	for x, y := range seq {
		return x, y, true, stop
	}

	var (
		zeroT T
		zeroU U
	)
	return zeroT, zeroU, false, func() {}
}
