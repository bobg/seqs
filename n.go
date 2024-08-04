package seqs

import "iter"

// FirstN produces a sequence containing the first n elements of the input
// (or all of the input, if there are fewer than n elements).
// Remaining elements of the input are not consumed.
func FirstN[T any](inp iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull(inp)
		defer stop()

		for n > 0 {
			val, ok := next()
			if !ok {
				return
			}
			if !yield(val) {
				return
			}
			n--
		}
	}
}

// FirstN2 produces a sequence containing the first n elements of the input
// (or all of the input, if there are fewer than n elements).
// Remaining elements of the input are not consumed.
func FirstN2[T, U any](inp iter.Seq2[T, U], n int) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		next, stop := iter.Pull2(inp)
		defer stop()

		for n > 0 {
			x, y, ok := next()
			if !ok {
				return
			}
			if !yield(x, y) {
				return
			}
			n--
		}
	}
}

// LastN produces a slice containing the last n elements of the input
// (or all of the input, if there are fewer than n elements).
// If the input is infinite, LastN will never return.
func LastN[T any, S ~func(func(T) bool)](inp S, n int) []T {
	var (
		buf   = make([]T, 0, n)
		start = 0
	)
	for val := range inp {
		if len(buf) < n {
			buf = append(buf, val)
			continue
		}
		buf[start] = val
		start = (start + 1) % n
	}
	rotateSlice(buf, -start)
	return buf
}

// SkipN copies the input iterator to the output,
// skipping the first N elements.
func SkipN[T any](inp iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull(inp)
		defer stop()

		for n > 0 {
			if _, ok := next(); !ok {
				return
			}
			n--
		}

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

// SkipN2 copies the input iterator to the output,
// skipping the first N elements.
func SkipN2[T, U any](inp iter.Seq2[T, U], n int) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		next, stop := iter.Pull2(inp)
		defer stop()

		for n > 0 {
			if _, _, ok := next(); !ok {
				return
			}
			n--
		}

		for {
			x, y, ok := next()
			if !ok {
				return
			}
			if !yield(x, y) {
				return
			}
		}
	}
}

// RotateSlice rotates a slice in place by n places to the right.
// (With negative n, it's to the left.)
// Example: RotateSlice([D, E, A, B, C], 3) -> [A, B, C, D, E]
func rotateSlice[T any, S ~[]T](s S, n int) {
	if n < 0 {
		// Convert left-rotation to right-rotation.
		n = -n
		n %= len(s)
		n = len(s) - n
	} else {
		n %= len(s)
	}
	if n == 0 {
		return
	}
	tmp := make([]T, n)
	copy(tmp, s[len(s)-n:])
	copy(s[n:], s)
	copy(s, tmp)
}
