package seqs

import "iter"

// A Zipped is a pair of zipped values, one of which may be missing,
// drawn from two different sequences.
type Zipped[V1, V2 any] struct {
	V1  V1
	Ok1 bool // whether V1 is present (if not, it will be zero)
	V2  V2
	Ok2 bool // whether V2 is present (if not, it will be zero)
}

// Zip returns an iterator that iterates x and y in parallel,
// yielding Zipped values of successive elements of x and y.
// If one sequence ends before the other, the iteration continues
// with Zipped values in which either Ok1 or Ok2 is false,
// depending on which sequence ended first.
//
// Zip is a useful building block for adapters that process
// pairs of sequences. For example, Equal can be defined as:
//
//	func Equal[V comparable](x, y iter.Seq[V]) bool {
//		for z := range Zip(x, y) {
//			if z.Ok1 != z.Ok2 || z.V1 != z.V2 {
//				return false
//			}
//		}
//		return true
//	}
func Zip[V1, V2 any](x iter.Seq[V1], y iter.Seq[V2]) iter.Seq[Zipped[V1, V2]] {
	return func(yield func(z Zipped[V1, V2]) bool) {
		next, stop := iter.Pull(y)
		defer stop()
		v2, ok2 := next()
		for v1 := range x {
			if !yield(Zipped[V1, V2]{v1, true, v2, ok2}) {
				return
			}
			v2, ok2 = next()
		}
		var zv1 V1
		for ok2 {
			if !yield(Zipped[V1, V2]{zv1, false, v2, ok2}) {
				return
			}
			v2, ok2 = next()
		}
	}
}

type Zipped2[K1, V1, K2, V2 any] struct {
	K1  K1
	V1  V1
	Ok1 bool // whether K1, V1 are present (if not, they will be zero)
	K2  K2
	V2  V2
	Ok2 bool // whether K2, V2 are present (if not, they will be zero)
}

// Zip2 returns an iterator that iterates x and y in parallel,
// yielding Zipped2 values of successive elements of x and y.
// If one sequence ends before the other, the iteration continues
// with Zipped2 values in which either Ok1 or Ok2 is false,
// depending on which sequence ended first.
//
// Zip2 is a useful building block for adapters that process
// pairs of sequences. For example, Equal2 can be defined as:
//
//	func Equal2[K, V comparable](x, y iter.Seq2[K, V]) bool {
//		for z := range Zip2(x, y) {
//			if z.Ok1 != z.Ok2 || z.K1 != z.K2 || z.V1 != z.V2 {
//				return false
//			}
//		}
//		return true
//	}
func Zip2[K1, V1, K2, V2 any](x iter.Seq2[K1, V1], y iter.Seq2[K2, V2]) iter.Seq[Zipped2[K1, V1, K2, V2]] {
	return func(yield func(z Zipped2[K1, V1, K2, V2]) bool) {
		next, stop := iter.Pull2(y)
		defer stop()
		k2, v2, ok2 := next()
		for k1, v1 := range x {
			if !yield(Zipped2[K1, V1, K2, V2]{k1, v1, true, k2, v2, ok2}) {
				return
			}
			k2, v2, ok2 = next()
		}
		var zk1 K1
		var zv1 V1
		for ok2 {
			if !yield(Zipped2[K1, V1, K2, V2]{zk1, zv1, false, k2, v2, ok2}) {
				return
			}
			k2, v2, ok2 = next()
		}
	}
}

// Zip takes two iterators and produces a new iterator containing pairs of corresponding elements.
// If one input iterator ends before the other,
// Zip produces zero values of the appropriate type in constructing pairs.
func ZipVals[T, U any](t iter.Seq[T], u iter.Seq[U]) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		for z := range Zip(t, u) {
			if !yield(z.V1, z.V2) {
				return
			}
		}
	}
}
