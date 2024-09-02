package seqs

import "iter"

// Empty is an empty sequence that can be used where an [iter.Seq] is expected.
// Usage note: you generally don't want to call this function,
// just refer to it as Empty[typename].
func Empty[T any](func(T) bool) {}

// Empty2 is an empty sequence that can be used where an [iter.Seq2] is expected.
// Usage note: you generally don't want to call this function,
// just refer to it as Empty2[typename1, typename2].
func Empty2[T, U any](func(T, U) bool) {}

// CheckEmpty tells whether the given iterator is empty.
// To determine whether it is, it consumes the first element of the iterator.
// It returns a copy of the original iterator
// (with any consumed element restored)
// and a boolean that is true for empty and false otherwise.
func CheckEmpty[T any](inp iter.Seq[T]) (iter.Seq[T], bool) {
	_, ok, seq := Peek(inp)
	return seq, !ok
}

// CheckEmpty2 tells whether the given iterator is empty.
// To determine whether it is, it consumes the first pair from the iterator.
// It returns a copy of the original iterator
// (with any consumed pair restored)
// and a boolean that is true for empty and false otherwise.
func CheckEmpty2[T, U any](inp iter.Seq2[T, U]) (iter.Seq2[T, U], bool) {
	_, _, ok, seq := Peek2(inp)
	return seq, !ok
}
