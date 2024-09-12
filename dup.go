package seqs

import (
	"iter"
	"sync"
)

// Dup duplicates the contents of an iterator,
// producing n new iterators,
// each containing the members of the original.
//
// An internal buffer grows to roughly the size
// of the difference between the output iterator that is farthest ahead in the stream,
// and the one that is farthest behind.
func Dup[T any](inp iter.Seq[T], n int) []iter.Seq[T] {
	var (
		ch = make(chan T)

		mu        sync.Mutex
		buf       []T
		bufOffset int
		offsets   = make([]int, n)
		result    []iter.Seq[T]
	)

	go func() { // xxx exit early if all output iterators are done
		for val := range inp {
			ch <- val
		}
		close(ch)
	}()

	for i := 0; i < n; i++ {
		helper := func(yield func(T) bool) bool {
			mu.Lock()
			defer mu.Unlock()

			bufEnd := bufOffset + len(buf)
			for offsets[i] >= bufEnd {
				val, ok := <-ch
				if !ok {
					return false
				}
				buf = append(buf, val)
				bufEnd++
			}
			val := buf[offsets[i]-bufOffset]
			offsets[i]++
			return yield(val)
		}
		result = append(result, func(yield func(T) bool) {
			for {
				if !helper(yield) {
					return
				}
			}
		})
	}

	return result
}
