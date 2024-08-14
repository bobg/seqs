package seqs

import "iter"

// Concat returns an iterator over the concatenation of the input sequences.
func Concat[T any](seqs ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, seq := range seqs {
			for e := range seq {
				if !yield(e) {
					return
				}
			}
		}
	}
}

// Concat2 returns an iterator over the concatenation of the sequences.
func Concat2[T, U any](seqs ...iter.Seq2[T, U]) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		for _, seq := range seqs {
			for k, v := range seq {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
