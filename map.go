package seqs

import "iter"

// Map returns an iterator over the values of inp transformed by the function f.
func Map[T, U any, F ~func(T) U](inp iter.Seq[T], f F) iter.Seq[U] {
	seq, _ := Mapx(inp, func(val T) (U, error) {
		return f(val), nil
	})
	return seq
}

// Mapx is the extended form of [Map].
// It produces an iterator of values transformed from an input iterator by a mapping function.
// If the mapping function returns an error,
// iteration stops and the error is available by dereferencing the returned pointer.
func Mapx[T, U any, F ~func(T) (U, error)](inp iter.Seq[T], f F) (iter.Seq[U], *error) {
	var err error

	g := func(yield func(U) bool) {
		for val := range inp {
			var out U

			out, err = f(val)
			if err != nil {
				return
			}
			if !yield(out) {
				return
			}
		}
	}

	return g, &err
}

// Map2 returns an iterator over the pairs of values of inp transformed by the function f.
func Map2[T1, U1, T2, U2 any, F ~func(T1, U1) (T2, U2)](inp iter.Seq2[T1, U1], f F) iter.Seq2[T2, U2] {
	seq, _ := Map2x(inp, func(t1 T1, u1 U1) (T2, U2, error) {
		t2, u2 := f(t1, u1)
		return t2, u2, nil
	})
	return seq
}

// Map2x is the extended form of [Map2].
// It produces an iterator of values transformed from an input iterator by a mapping function.
// If the mapping function returns an error,
// iteration stops and the error is available by dereferencing the returned pointer.
func Map2x[T1, U1, T2, U2 any, F ~func(T1, U1) (T2, U2, error)](inp iter.Seq2[T1, U1], f F) (iter.Seq2[T2, U2], *error) {
	var err error

	g := func(yield func(T2, U2) bool) {
		for k, v := range inp {
			var out1 T2
			var out2 U2

			out1, out2, err = f(k, v)
			if err != nil {
				return
			}
			if !yield(out1, out2) {
				return
			}
		}
	}

	return g, &err
}
