package seqs

import (
	"cmp"
	"iter"
	"slices"
)

// Merge merges values from multiple sequences ordered according to [cmp.Less]
// into a new sequence that is also ordered according to [cmp.Less].
func Merge[T cmp.Ordered](inps ...iter.Seq[T]) iter.Seq[T] {
	return MergeFunc(inps, cmp.Less)
}

// MergeFunc merges values from multiple sequences ordered according to the given function,
// which should return true iff the first argument is "less"
// (i.e., sorts earlier than)
// than the second.
func MergeFunc[T any](inps []iter.Seq[T], less func(T, T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		var (
			nexts     = make([]func() (T, bool), len(inps))
			vals      = make([]T, len(inps))
			done      = make([]bool, len(inps))
			remaining = len(inps)
		)

		for i, inp := range inps {
			next, stop := iter.Pull(inp)
			defer stop()

			nexts[i] = next
			val, ok := next()
			if ok {
				vals[i] = val
			} else {
				done[i] = true
				remaining--
			}
		}

		for remaining > 1 {
			var (
				first    = true
				minVal   T
				minIndex int
			)
			for i := 0; i < len(inps); i++ {
				if done[i] {
					continue
				}
				if first {
					minVal = vals[i]
					minIndex = i
					first = false
					continue
				}
				if less(vals[i], minVal) {
					minVal = vals[i]
					minIndex = i
				}
			}

			if !yield(minVal) {
				return
			}

			if val, ok := nexts[minIndex](); ok {
				vals[minIndex] = val
			} else {
				done[minIndex] = true
				remaining--
			}
		}

		// Only one input sequence remaining.
		// Find it and copy its values to the output.

		index := slices.Index(done, false)
		if !yield(vals[index]) {
			return
		}

		next := nexts[index]
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
