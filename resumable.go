package seqs

import (
	"iter"
	"sync/atomic"
)

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
