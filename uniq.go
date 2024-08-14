package seqs

import "iter"

// Uniq filters out duplicate values from a sequence,
// yielding only the first occurrence of each value.
//
// Only adjacent duplicates are removed.
// For example, {1, 1, 2, 2, 2, 3} becomes {1, 2, 3},
// but {1, 2, 1, 2, 3} remains {1, 2, 1, 2, 3}.
func Uniq[T comparable](inp iter.Seq[T]) iter.Seq[T] {
	return UniqFunc(inp, func(a, b T) bool { return a == b })
}

// UniqFunc filters out duplicate values from a sequence.
// Two values a and b are considered duplicates if f(a, b) returns true.
//
// Only adjacent duplicates are removed.
func UniqFunc[T any, F ~func(T, T) bool](inp iter.Seq[T], f F) iter.Seq[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull(inp)
		defer stop()

		prev, ok := next()
		if !ok {
			return
		}
		if !yield(prev) {
			return
		}
		for {
			val, ok := next()
			if !ok {
				return
			}
			if f(prev, val) {
				continue
			}
			if !yield(val) {
				return
			}
			prev = val
		}
	}
}
