package seqs

import (
	"fmt"
	"iter"
)

// Pages converts an iterator of items into an iterator of pages of items.
// Each page is a slice of up to pageSize items.
func Pages[T any](inp iter.Seq[T], pageSize int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		page := make([]T, 0, pageSize)

		for x := range inp {
			page = append(page, x)
			if len(page) >= pageSize {
				if !yield(page) {
					return
				}
				page = make([]T, 0, pageSize)
			}
		}

		if len(page) > 0 {
			yield(page)
		}
	}
}

// FromPages produces an iterator of items from a function that returns items a page at a time.
// The boolean result from nextpage indicates whether there are more pages to come.
func FromPages[T any](nextpage func() ([]T, bool)) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			page, ok := nextpage()
			fmt.Printf("xxx got a page and %v\n", ok)
			if !ok {
				return
			}
			for _, item := range page {
				if !yield(item) {
					return
				}
			}
		}
	}
}
