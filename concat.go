package seqs

import "iter"

// Concat returns an iterator over the concatenation of the sequences.
func Concat[V any](seqs ...iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
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
func Concat2[K, V any](seqs ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, seq := range seqs {
			for k, v := range seq {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
