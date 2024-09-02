package seqs

import "iter"

// Peek returns the first value in the given iterator.
// If the iterator is empty, the returned boolean is false, otherwise it's true.
// The first value of the input iterator, if there is one, is consumed.
// The returned iterator is a copy of the original iterator with any consumed value restored.
func Peek[T any](inp iter.Seq[T]) (T, bool, iter.Seq[T]) {
	next, stop := iter.Pull(inp)

	if v, ok := next(); ok {
		return v, true, func(yield func(T) bool) {
			defer stop()

			if !yield(v) {
				return
			}
			for {
				v, ok := next()
				if !ok {
					return
				}
				if !yield(v) {
					return
				}
			}
		}
	}

	stop()

	var zero T
	return zero, false, Empty[T]
}

// Peek2 returns the first pair of values in the given iterator.
// If the iterator is empty, the returned boolean is false, otherwise it's true.
// The first pair of the input iterator, if there is one, is consumed.
// The returned iterator is a copy of the original iterator with any consumed pair restored.
func Peek2[T, U any](inp iter.Seq2[T, U]) (T, U, bool, iter.Seq2[T, U]) {
	next, stop := iter.Pull2(inp)

	if t, u, ok := next(); ok {
		return t, u, true, func(yield func(T, U) bool) {
			defer stop()

			if !yield(t, u) {
				return
			}
			for {
				t, u, ok := next()
				if !ok {
					return
				}
				if !yield(t, u) {
					return
				}
			}
		}
	}

	stop()

	var (
		zeroT T
		zeroU U
	)
	return zeroT, zeroU, false, Empty2[T, U]
}
