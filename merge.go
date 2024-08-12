package seqs

import (
	"cmp"
	"iter"
	"slices"
)

// Merge merges two sequences of ordered values.
// Values appear in the output once for each time they appear in x
// and once for each time they appear in y.
// If the two input sequences are not ordered,
// the output sequence will not be ordered,
// but it will still contain every value from x and y exactly once.
//
// Merge is equivalent to calling MergeFunc with cmp.Compare[T]
// as the ordering function.
func Merge[T cmp.Ordered](x, y iter.Seq[T]) iter.Seq[T] {
	return MergeFunc(x, y, cmp.Compare[T])
}

// MergeFunc merges two sequences of values ordered by the function f.
// Values appear in the output once for each time they appear in x
// and once for each time they appear in y.
// When equal values appear in both sequences,
// the output contains the values from x before the values from y.
// If the two input sequences are not ordered by f,
// the output sequence will not be ordered by f,
// but it will still contain every value from x and y exactly once.
func MergeFunc[T any](x, y iter.Seq[T], f func(T, T) int) iter.Seq[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull(y)
		defer stop()
		v2, ok2 := next()
		for v1 := range x {
			for ok2 && f(v1, v2) > 0 {
				if !yield(v2) {
					return
				}
				v2, ok2 = next()
			}
			if !yield(v1) {
				return
			}
		}
		for ok2 {
			if !yield(v2) {
				return
			}
			v2, ok2 = next()
		}
	}
}

// Merge2 merges two sequences of key-value pairs ordered by their keys.
// Pairs appear in the output once for each time they appear in x
// and once for each time they appear in y.
// If the two input sequences are not ordered by their keys,
// the output sequence will not be ordered by its keys,
// but it will still contain every pair from x and y exactly once.
//
// Merge2 is equivalent to calling MergeFunc2 with cmp.Compare[K]
// as the ordering function.
func Merge2[K cmp.Ordered, V any](x, y iter.Seq2[K, V]) iter.Seq2[K, V] {
	return MergeFunc2(x, y, cmp.Compare[K])
}

// MergeFunc2 merges two sequences of key-value pairs ordered by the function f.
// Pairs appear in the output once for each time they appear in x
// and once for each time they appear in y.
// When pairs with equal keys appear in both sequences,
// the output contains the pairs from x before the pairs from y.
// If the two input sequences are not ordered by f,
// the output sequence will not be ordered by f,
// but it will still contain every pair from x and y exactly once.
func MergeFunc2[K, V any](x, y iter.Seq2[K, V], f func(K, K) int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		next, stop := iter.Pull2(y)
		defer stop()
		k2, v2, ok2 := next()
		for k1, v1 := range x {
			for ok2 && f(k1, k2) > 0 {
				if !yield(k2, v2) {
					return
				}
				k2, v2, ok2 = next()
			}
			if !yield(k1, v1) {
				return
			}
		}
		for ok2 {
			if !yield(k2, v2) {
				return
			}
			k2, v2, ok2 = next()
		}
	}
}

func MergeAll[T cmp.Ordered](seqs ...iter.Seq[T]) iter.Seq[T] {
	return MergeAllFunc(seqs, cmp.Compare[T])
}

func MergeAllFunc[T cmp.Ordered](inps []iter.Seq[T], f func(T, T) int) iter.Seq[T] {
	switch len(inps) {
	case 0:
		return Empty[T]
	case 1:
		return inps[0]
	case 2:
		return MergeFunc(inps[0], inps[1], f)
	default:
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
					if f(vals[i], minVal) < 0 {
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
}
