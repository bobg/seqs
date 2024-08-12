package seqs

import "iter"

// Last returns the last element of the input sequence and true.
// If the input is empty, Last returns the zero value of T instead, and false.
// Last consumes the entire input sequence.
// Beware of infinite input sequences!
func Last[T any](inp iter.Seq[T]) (T, bool) {
	var (
		last T
		ok   bool
	)

	for val := range inp {
		last = val
		ok = true
	}
	return last, ok
}

// Last2 returns the last pair of elements from the pairwise input sequence, and true.
// If the input is empty, Last2 returns the zero values of T and U instead, and false.
// Last2 consumes the entire input sequence.
// Beware of infinite input sequences!
func Last2[T, U any](inp iter.Seq2[T, U]) (lastT T, lastU U, ok bool) {
	for x, y := range inp {
		lastT, lastU = x, y
		ok = true
	}
	return lastT, lastU, ok
}

// LastN produces a slice containing the last n elements of the input
// (or all of the input, if there are fewer than n elements).
// Beware of infinite input sequences!
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
