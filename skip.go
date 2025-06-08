package seqs

import "iter"

// SkipUntil copies the input iterator to the output,
// discarding the initial elements until the first one that causes f to return true.
// It is the same as SkipWhile(inp, notF),
// where notF is the boolean inverse of f.
func SkipUntil[T any](inp iter.Seq[T], f func(T) bool) iter.Seq[T] {
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

// SkipWhile copies the input iterator to the output,
// discarding the initial elements while f returns true.
// It is the same as SkipUntil(inp, notF),
// where notF is the boolean inverse of f.
func SkipWhile[T any](inp iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return SkipUntil(inp, func(val T) bool { return !f(val) })
}

// SkipN copies the input iterator to the output,
// skipping the first N elements.
func SkipN[T any](inp iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		skipping := n > 0
		for val := range inp {
			if skipping {
				n--
				if n == 0 {
					skipping = false
				}
				continue
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
		skipping := n > 0
		for k, v := range inp {
			if skipping {
				n--
				if n == 0 {
					skipping = false
				}
				continue
			}
			if !yield(k, v) {
				return
			}
		}
	}
}
