package seqs

import (
	"cmp"
	"iter"
)

// CommLeft takes two ordered sequences and produces a sequence of elements that are in left but not in right.
// For example, if left is [1, 2, 3, 4] and right is [2, 4, 6, 8], CommLeft produces [1, 3].
func CommLeft[T cmp.Ordered](left, right iter.Seq[T]) iter.Seq[T] {
	return CommLeftFunc(left, right, cmp.Compare)
}

// CommLeftFunc is like [CommLeft], but it uses a custom comparison function.
// The function cmp must return a negative value if its arguments are properly ordered,
// a positive value if they are reversed,
// and zero if they are equal.
func CommLeftFunc[T any](left, right iter.Seq[T], cmp func(T, T) int) iter.Seq[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull(right)
		defer stop()

		r, ok := next()

		for l := range left {

			if ok {
				c := cmp(l, r)
				for ok && c > 0 {
					r, ok = next()
					if ok {
						c = cmp(l, r)
					}
				}
				if ok && c == 0 {
					r, ok = next()
					continue
				}
			}
			if !yield(l) {
				return
			}
		}
	}
}

// CommRight takes two ordered sequences and produces a sequence of elements that are in right but not in left.
// For example, if left is [1, 2, 3, 4] and right is [2, 4, 6, 8], CommRight produces [6, 8].
func CommRight[T cmp.Ordered](left, right iter.Seq[T]) iter.Seq[T] {
	return CommRightFunc(left, right, cmp.Compare)
}

// CommRightFunc is like [CommRight], but it uses a custom comparison function.
// The function cmp must return a negative value if its arguments are properly ordered,
// a positive value if they are reversed,
// and zero if they are equal.
func CommRightFunc[T any](left, right iter.Seq[T], cmp func(T, T) int) iter.Seq[T] {
	return CommLeftFunc(right, left, cmp)
}

// CommBoth takes two ordered sequences and produces a sequence of elements that are in both left and right.
// For example, if left is [1, 2, 3, 4] and right is [2, 4, 6, 8], CommBoth produces [2, 4].
func CommBoth[T cmp.Ordered](left, right iter.Seq[T]) iter.Seq[T] {
	return CommBothFunc(left, right, cmp.Compare)
}

// CommBothFunc is like [CommBoth], but it uses a custom comparison function.
// The function cmp must return a negative value if its arguments are properly ordered,
// a positive value if they are reversed,
// and zero if they are equal.
func CommBothFunc[T any](left, right iter.Seq[T], cmp func(T, T) int) iter.Seq[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull(right)
		defer stop()

		r, ok := next()

		for l := range left {
			if !ok {
				return
			}
			c := cmp(l, r)
			for ok && c > 0 {
				r, ok = next()
				if ok {
					c = cmp(l, r)
				}
			}
			if ok && c == 0 {
				if !yield(l) {
					return
				}
				r, ok = next()
			}
		}
	}
}
