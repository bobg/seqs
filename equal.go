package seqs

import "iter"

// Equal reports whether the two sequences are equal.
func Equal[V comparable](x, y iter.Seq[V]) bool {
	for z := range Zip(x, y) {
		if z.Ok1 != z.Ok2 || z.V1 != z.V2 {
			return false
		}
	}
	return true
}

// Equal2 reports whether the two sequences are equal.
func Equal2[K, V comparable](x, y iter.Seq2[K, V]) bool {
	for z := range Zip2(x, y) {
		if z.Ok1 != z.Ok2 || z.K1 != z.K2 || z.V1 != z.V2 {
			return false
		}
	}
	return true
}

// EqualFunc reports whether the two sequences are equal according to the function f.
func EqualFunc[V1, V2 any](x iter.Seq[V1], y iter.Seq[V2], f func(V1, V2) bool) bool {
	for z := range Zip(x, y) {
		if z.Ok1 != z.Ok2 || !f(z.V1, z.V2) {
			return false
		}
	}
	return true
}

// EqualFunc2 reports whether the two sequences are equal according to the function f.
func EqualFunc2[K1, V1, K2, V2 any](x iter.Seq2[K1, V1], y iter.Seq2[K2, V2], f func(K1, V1, K2, V2) bool) bool {
	for z := range Zip2(x, y) {
		if z.Ok1 != z.Ok2 || !f(z.K1, z.V1, z.K2, z.V2) {
			return false
		}
	}
	return true
}
