package seqs

import "iter"

// Accum repeatedly applies a function to the elements of an iterator
// and produces an iterator over the accumulated values.
// On each call, the function receives the accumulated value and the next element of the input iterator.
// The first call receives the initial value as the accumulated value.
func Accum[T, A any, F ~func(A, T) A](inp iter.Seq[T], init A, f F) iter.Seq[A] {
	seq, _ := Accumx(inp, init, func(acc A, t T) (A, error) {
		return f(acc, t), nil
	})
	return seq
}

// Accumx repeatedly applies a function to the elements of an iterator
// and produces an iterator over the accumulated values,
// plus a pointer to an error.
// On each call, the function receives the accumulated value and the next element of the input iterator.
// The first call receives the initial value as the accumulated value.
// If the function returns an error, Accumx stops.
// The error pointer that Accumx returns may be dereferenced to discover that error,
// but only after the output iterator is fully consumed.
func Accumx[T, A any, F ~func(A, T) (A, error)](inp iter.Seq[T], init A, f F) (iter.Seq[A], *error) {
	var err error

	seq := func(yield func(A) bool) {
		acc := init
		for val := range inp {
			acc, err = f(acc, val)
			if err != nil || !yield(acc) {
				return
			}
		}
	}

	return seq, &err
}

// Accum2 repeatedly applies a function to the elements of a pairwise iterator
// and produces an iterator over the accumulated values.
// On each call, the function receives the accumulated value and the next pair of elements from the input iterator.
// The first call receives the initial value as the accumulated value.
func Accum2[T, U, A any, F ~func(A, T, U) A](inp iter.Seq2[T, U], init A, f F) iter.Seq[A] {
	seq, _ := Accum2x(inp, init, func(acc A, t T, u U) (A, error) {
		return f(acc, t, u), nil
	})
	return seq
}

// Accum2x repeatedly applies a function to the elements of an iterator
// and produces an iterator over the accumulated values,
// plus a pointer to an error.
// On each call, the function receives the accumulated value and the next pair of elements from the input iterator.
// The first call receives the initial value as the accumulated value.
// If the function returns an error, Accum2x stops.
// The error pointer that Accum2x returns may be dereferenced to discover that error,
// but only after the output iterator is fully consumed.
func Accum2x[T, U, A any, F ~func(A, T, U) (A, error)](inp iter.Seq2[T, U], init A, f F) (iter.Seq[A], *error) {
	var err error

	seq := func(yield func(A) bool) {
		acc := init
		for t, u := range inp {
			acc, err = f(acc, t, u)
			if err != nil || !yield(acc) {
				return
			}
		}
	}

	return seq, &err
}
