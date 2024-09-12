package seqs

import (
	"iter"
	"sync/atomic"
)

// Resumable turns an iterator into a resumable one.
// Normally, if you use an iterator in a "for ... range" loop,
// and break out of the loop early,
// you cannot resume the iteration later:
// a new "for ... range" loop with the same iterator will start a new iteration.
// But a resumable iterator will remember where it left off.
//
// In addition to the resumable iterator,
// this function returns a "stop" function
// that must be called to release resources
// when you're done with the iterator.
func Resumable[T any](inp iter.Seq[T]) (iter.Seq[T], func()) {
	next, stop := iter.Pull(inp)

	var stopped atomic.Bool

	result := func(yield func(T) bool) {
		for {
			val, ok := next()
			if !ok {
				stop()
				return
			}
			if !yield(val) {
				// Don't call stop here.
				// The caller may want to resume the sequence later,
				// and will then have to call stop itself.
				return
			}
		}
	}

	return result, func() {
		if stopped.CompareAndSwap(false, true) {
			stop()
		}
	}
}

// Resumable2 turns an iterator into a resumable one.
// Normally, if you use an iterator in a "for ... range" loop,
// and break out of the loop early,
// you cannot resume the iteration later:
// a new "for ... range" loop with the same iterator will start a new iteration.
// But a resumable iterator will remember where it left off.
//
// In addition to the resumable iterator,
// this function returns a "stop" function
// that must be called to release resources
// when you're done with the iterator.
func Resumable2[T, U any](inp iter.Seq2[T, U]) (iter.Seq2[T, U], func()) {
	next, stop := iter.Pull2(inp)

	var stopped atomic.Bool

	result := func(yield func(T, U) bool) {
		for {
			x, y, ok := next()
			if !ok {
				stop()
				return
			}
			if !yield(x, y) {
				// Don't call stop here.
				// The caller may want to resume the sequence later,
				// and will then have to call stop itself.
				return
			}
		}
	}

	return result, func() {
		if stopped.CompareAndSwap(false, true) {
			stop()
		}
	}
}
