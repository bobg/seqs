package seqs

import "iter"

// Peek returns the first value in the given iterator.
// If the iterator is empty, the returned boolean is false, otherwise it's true.
// The first value of the input iterator, if there is one, is consumed.
// The returned iterator is a copy of the original iterator with any consumed value restored.
//
// Note: if the caller fails to consume the returned iterator,
// this function can leak resources.
// One good way to make sure this does not happen,
// if the iterator is unneeded,
// is to pass the iterator to [Drain].
// Another is to call [iter.Pull] on it
// and immediately call the returned "stop" function.
//
// Deprecated: Ensuring this does not leak resources is not as ergonomic as it can be. Use [Peeker] instead.
func Peek[T any](inp iter.Seq[T]) (T, bool, iter.Seq[T]) {
	next, stop := iter.Pull(inp)

	if v, ok := next(); ok {
		return v, true, func(yield func(T) bool) {
			defer stop()

			if !yield(v) {
				return
			}
			for {
				v, ok := next()
				if !ok {
					return
				}
				if !yield(v) {
					return
				}
			}
		}
	}

	stop()

	var zero T
	return zero, false, Empty[T]
}

// Peeker is like [iter.Pull] but adds the ability to "peek" at the next value in a sequence without consuming it.
// As with iter.Pull, callers must be sure to call stop (usually via defer)
// when finished with the resulting "peeker."
func Peeker[T any](inp iter.Seq[T]) (next, peek func() (T, bool), stop func()) {
	next1, stop := iter.Pull(inp)

	var peeked *T

	next = func() (T, bool) {
		if peeked != nil {
			result := *peeked
			peeked = nil
			return result, true
		}
		return next1()
	}

	peek = func() (T, bool) {
		if peeked != nil {
			return *peeked, true
		}
		val, ok := next1()
		if !ok {
			return val, false
		}
		peeked = &val
		return val, true
	}

	return next, peek, stop
}

// Peek2 returns the first pair of values in the given iterator.
// If the iterator is empty, the returned boolean is false, otherwise it's true.
// The first pair of the input iterator, if there is one, is consumed.
// The returned iterator is a copy of the original iterator with any consumed pair restored.
//
// Note: if the caller fails to consume the returned iterator,
// this function can leak resources.
// One good way to make sure this does not happen,
// if the iterator is unneeded,
// is to pass the iterator to [Drain2].
// Another is to call [iter.Pull2] on it
// and immediately call the returned "stop" function.
//
// Deprecated: Ensuring this does not leak resources is not as ergonomic as it can be. Use [Peeker2] instead.
func Peek2[T, U any](inp iter.Seq2[T, U]) (T, U, bool, iter.Seq2[T, U]) {
	next, stop := iter.Pull2(inp)

	if t, u, ok := next(); ok {
		return t, u, true, func(yield func(T, U) bool) {
			defer stop()

			if !yield(t, u) {
				return
			}
			for {
				t, u, ok := next()
				if !ok {
					return
				}
				if !yield(t, u) {
					return
				}
			}
		}
	}

	stop()

	var (
		zeroT T
		zeroU U
	)
	return zeroT, zeroU, false, Empty2[T, U]
}

// Peeker2 is like [iter.Pull2] but adds the ability to "peek" at the next pair of values in a sequence without consuming them.
// As with iter.Pull2, callers must be sure to call stop (usually via defer)
// when finished with the resulting "peeker."
func Peeker2[T, U any](inp iter.Seq2[T, U]) (next, peek func() (T, U, bool), stop func()) {
	next1, stop := iter.Pull2(inp)

	var (
		peekedT *T
		peekedU *U
	)

	next = func() (T, U, bool) {
		if peekedT != nil {
			resultT, resultU := *peekedT, *peekedU
			peekedT, peekedU = nil, nil
			return resultT, resultU, true
		}
		return next1()
	}

	peek = func() (T, U, bool) {
		if peekedT != nil {
			return *peekedT, *peekedU, true
		}
		valT, valU, ok := next1()
		if !ok {
			return valT, valU, false
		}
		peekedT, peekedU = &valT, &valU
		return valT, valU, true
	}

	return next, peek, stop
}
