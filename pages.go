package seqs

import "iter"

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
