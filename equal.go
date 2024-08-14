package seqs

import "iter"

// Equal reports whether the two sequences are equal.
// It consumes elements from both sequences until one of them is exhausted,
// or until it finds a pair of elements that are not equal.
// If the sequences are equal and infinite, Equal will not terminate.
func Equal[T comparable](x, y iter.Seq[T]) bool {
	for z := range Zip(x, y) {
		if z.Ok1 != z.Ok2 || z.V1 != z.V2 {
			return false
		}
	}
	return true
}

// Equal2 reports whether the two sequences are equal.
// It consumes pairs of elements from both sequences until one of them is exhausted,
// or until it finds a pair of element-pairs that are not equal.
// If the sequences are equal and infinite, Equal2 will not terminate.
func Equal2[T, U comparable](x, y iter.Seq2[T, U]) bool {
	for z := range Zip2(x, y) {
		if z.Ok1 != z.Ok2 || z.K1 != z.K2 || z.V1 != z.V2 {
			return false
		}
	}
	return true
}

// EqualFunc reports whether the two sequences are equal according to the function f.
// It consumes elements from both sequences until one of them is exhausted,
// or until it finds a pair of elements that are not equal.
// If the sequences are equal and infinite, EqualFunc will not terminate.
func EqualFunc[T, U any](x iter.Seq[T], y iter.Seq[U], f func(T, U) bool) bool {
	for z := range Zip(x, y) {
		if z.Ok1 != z.Ok2 || !f(z.V1, z.V2) {
			return false
		}
	}
	return true
}

// EqualFunc2 reports whether the two sequences are equal according to the function f.
// It consumes pairs of elements from both sequences until one of them is exhausted,
// or until it finds a pair of element-pairs that are not equal.
// If the sequences are equal and infinite, EqualFunc2 will not terminate.
func EqualFunc2[K1, V1, K2, V2 any](x iter.Seq2[K1, V1], y iter.Seq2[K2, V2], f func(K1, V1, K2, V2) bool) bool {
	for z := range Zip2(x, y) {
		if z.Ok1 != z.Ok2 || !f(z.K1, z.V1, z.K2, z.V2) {
			return false
		}
	}
	return true
}
