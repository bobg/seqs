package seqs

import (
	"iter"

	"github.com/bobg/go-generics/v3/slices"
)

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
// LastN consumes the entire input sequence.
// Beware of infinite input sequences!
func LastN[T any](inp iter.Seq[T], n int) []T {
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
	slices.Rotate(buf, -start)
	return buf
}

// LastN2 produces a slice containing the last n pairs of elements of the input
// (or all of the input, if there are fewer than n pairs).
// LastN2 consumes the entire input sequence.
// Beware of infinite input sequences!
func LastN2[T, U any](inp iter.Seq2[T, U], n int) []Pair[T, U] {
	var (
		buf   = make([]Pair[T, U], 0, n)
		start = 0
	)
	for x, y := range inp {
		if len(buf) < n {
			buf = append(buf, Pair[T, U]{x, y})
			continue
		}
		buf[start] = Pair[T, U]{x, y}
		start = (start + 1) % n
	}
	slices.Rotate(buf, -start)
	return buf
}
