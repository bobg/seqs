package seqs

import "iter"

// Peek returns the first value in the given iterator.
// If the iterator is empty, the returned boolean is false, otherwise it's true.
// The first value of the input iterator, if there is one, is consumed.
// The returned iterator is a copy of the original iterator with any consumed value restored.
// Also returned is a "stop" function that must be used when finished with the returned iterator.
func Peek[T any](inp iter.Seq[T]) (T, bool, iter.Seq[T], func()) {
	var stop func()
	inp, stop = Resumable(inp)

	for v := range inp {
		out := func(yield func(T) bool) {
			if !yield(v) {
				return
			}
			for v := range inp {
				if !yield(v) {
					return
				}
			}
		}
		return v, true, out, stop
	}

	var zero T
	return zero, false, inp, stop
}

// Peek2 returns the first pair of values in the given iterator.
// If the iterator is empty, the returned boolean is false, otherwise it's true.
// The first pair of the input iterator, if there is one, is consumed.
// The returned iterator is a copy of the original iterator with any consumed pair restored.
// Also returned is a "stop" function that must be used when finished with the returned iterator.
func Peek2[T, U any](inp iter.Seq2[T, U]) (T, U, bool, iter.Seq2[T, U], func()) {
	var stop func()
	inp, stop = Resumable2(inp)

	for x, y := range inp {
		out := func(yield func(T, U) bool) {
			if !yield(x, y) {
				return
			}
			for x, y := range inp {
				if !yield(x, y) {
					return
				}
			}
		}
		return x, y, true, out, stop
	}

	var zeroT T
	var zeroU U
	return zeroT, zeroU, false, inp, stop
}
