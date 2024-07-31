package seqs

import "iter"

// FirstN produces an iterator containing the first n elements of the input
// (or all of the input, if there are fewer than n elements).
// Remaining elements of the input are not consumed.
// It is the caller's responsibility to release any associated resources.
func FirstN[T any](inp iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i, val := range Enumerate(inp) {
			if i >= n {
				return
			}
			if !yield(val) {
				return
			}
		}
	}
}

// LastN produces a slice containing the last n elements of the input iterator
// (or all of the input, if there are fewer than n elements).
// There is no guarantee that any elements will ever be produced:
// the input iterator may be infinite!
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
		for i, val := range Enumerate(inp) {
			if i < n {
				continue
			}
			if !yield(val) {
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
