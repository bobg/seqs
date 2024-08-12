package seqs

import "iter"

// SkipUntil copies the input iterator to the output,
// discarding the initial elements until the first one that causes f to return true.
// That element and the remaining elements of inp are included in the output,
// and f is not called again.
func SkipUntil[T any, F ~func(T) bool](inp iter.Seq[T], f F) iter.Seq[T] {
	skipping := true
	return Filter(inp, func(val T) bool {
		if !skipping {
			return true
		}
		if f(val) {
			skipping = false
		}
		return !skipping
	})
}

// SkipN copies the input iterator to the output,
// skipping the first N elements.
func SkipN[T any](inp iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull(inp)
		defer stop()

		for n > 0 {
			if _, ok := next(); !ok {
				return
			}
			n--
		}

		for {
			val, ok := next()
			if !ok {
				return
			}
			if !yield(val) {
				return
			}
		}
	}
}

// SkipN2 copies the input iterator to the output,
// skipping the first N elements.
func SkipN2[T, U any](inp iter.Seq2[T, U], n int) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		next, stop := iter.Pull2(inp)
		defer stop()

		for n > 0 {
			if _, _, ok := next(); !ok {
				return
			}
			n--
		}

		for {
			x, y, ok := next()
			if !ok {
				return
			}
			if !yield(x, y) {
				return
			}
		}
	}
}
