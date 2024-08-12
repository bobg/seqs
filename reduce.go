package seqs

import "iter"

// Reduce combines the values in seq using f.
// The result begins as init.
// Then for each value v in seq,
// it updates the result to be f(result, v).
// It returns the final result value.
//
// For example, if iterating over seq yields v1, v2, v3,
// Reduce returns f(f(f(init, v1), v2), v3).
//
// Reduce is equivalent to calling Last(Accum(seq, init, f)).
// It consumes the entire input sequence.
// Beware of infinite input!
func Reduce[T, A any, F ~func(A, T) A](seq iter.Seq[T], init A, f F) A {
	val, _ := Reducex(seq, init, func(acc A, t T) (A, error) {
		return f(acc, t), nil
	})
	return val
}

// Reducex combines the values in seq using f.
// The result begins as init.
// Then for each value v in seq,
// it updates the result to be f(result, v).
// It returns the final result value.
// If f returns an error,
// iteration stops and Reducex returns that error.
//
// Reducex consumes the entire input sequence.
// Beware of infinite input!
func Reducex[T, A any, F ~func(A, T) (A, error)](seq iter.Seq[T], init A, f F) (A, error) {
	acc, errptr := Accumx(seq, init, f)
	if last, ok := Last(acc); ok {
		return last, *errptr
	}
	return init, *errptr
}

// Reduce2 combines the values in seq using f.
// The result begins as init.
// Then for each value pair t,u in seq,
// it updates the result to be f(result, t, u).
// It returns the final result value.
//
// For example, if iterating over seq yields pairs t1,u1, t2,u2, t3,u3,
// Reduce2 returns f(f(f(init, t1, u1), t2, u2), t3, u3).
//
// Reduce2 is equivalent to calling Last2(Accum2(seq, init, f)).
// It consumes the entire input sequence.
// Beware of infinite input!
func Reduce2[T, U, A any, F ~func(A, T, U) A](seq iter.Seq2[T, U], init A, f F) A {
	acc, _ := Reduce2x(seq, init, func(acc A, t T, u U) (A, error) {
		return f(acc, t, u), nil
	})
	return acc
}

// Reduce2x combines the values in seq using f.
// The result begins as init.
// Then for each value pair t,u in seq,
// it updates the result to be f(result, t, u).
// It returns the final result value.
// If f returns an error,
// iteration stops and Reduce2x returns that error.
//
// Reduce2x consumes the entire input sequence.
// Beware of infinite input!
func Reduce2x[T, U, A any, F ~func(A, T, U) (A, error)](inp iter.Seq2[T, U], init A, f F) (A, error) {
	seq, errptr := Accum2x(inp, init, f)
	if last, ok := Last(seq); ok {
		return last, *errptr
	}
	return init, *errptr
}
