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
//
// The function f should return zero if its arguments are equal,
// a negative value if the arguments are in the proper order,
// and a positive value if they are in the reverse order.
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
//
// The function f should return zero if its arguments are equal,
// a negative value if the arguments are in the proper order,
// and a positive value if they are in the reverse order.
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

// MergeAll merges zero or more ordered sequences into a single one.
func MergeAll[T cmp.Ordered](seqs ...iter.Seq[T]) iter.Seq[T] {
	return MergeAllFunc(seqs, cmp.Compare[T])
}

// MergeAllFunc merges zero or more ordered sequences into a single one,
// using the ordering function f,
// which should return zero if its arguments are equal,
// a negative value if the arguments are in the proper order,
// and a positive value if they are in the reverse order.
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
				nexts = make([]func() (T, bool), 0, len(inps))
				vals  = make([]T, 0, len(inps))
			)

			for _, inp := range inps {
				next, stop := iter.Pull(inp)
				defer stop()

				val, ok := next()
				if ok {
					vals = append(vals, val)
					nexts = append(nexts, next)
				}
			}

			for len(vals) > 0 {
				var (
					best      = vals[0]
					bestIndex = 0
				)
				for i := 1; i < len(vals); i++ {
					val := vals[i]
					if f(val, best) < 0 {
						best = val
						bestIndex = i
					}
				}

				if !yield(best) {
					return
				}

				// Replace the best value with the next value from its input.
				nextval, ok := nexts[bestIndex]()
				if ok {
					vals[bestIndex] = nextval
				} else {
					// That input is exhausted.
					nexts = slices.Delete(nexts, bestIndex, bestIndex+1)
					vals = slices.Delete(vals, bestIndex, bestIndex+1)
				}
			}
		}
	}
}
