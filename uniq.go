package seqs

import "iter"

// Uniq filters out duplicate values from a sequence,
// yielding only the first occurrence of each value.
//
// Only adjacent duplicates are removed.
// For example, {1, 1, 2, 2, 2, 3} becomes {1, 2, 3},
// but {1, 2, 1, 1, 2, 2, 3} becomes {1, 2, 1, 2, 3}.
func Uniq[T comparable](inp iter.Seq[T]) iter.Seq[T] {
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
			if val == prev {
				continue
			}
			if !yield(val) {
				return
			}
			prev = val
		}
	}
}
