package seqs

import "iter"

// Err converts an error-returning function into a Go iterator
// (i.e., an iter.Seq[T])
// plus an error pointer.
//
// The function being converted must be a func(func(T) bool) error.
// (An iter.Seq[T] is a func(func(T) bool), without the error result.)
//
// The resulting error pointer is never nil,
// but the error it points to is not guaranteed to be populated
// until after the iterator has been fully consumed.
func Err[T any](f func(yield func(T) bool) error) (iter.Seq[T], *error) {
	var err error
	result := func(yield func(T) bool) {
		err = f(yield)
	}
	return result, &err
}

// Err2 converts an error-returning function into a Go iterator
// (i.e., an iter.Seq2[T, U])
// plus an error pointer.
//
// The function being converted must be a func(func(T, U) bool) error.
// (An iter.Seq2[T, U] is a func(func(T, U) bool), without the error result.)
//
// The resulting error pointer is never nil,
// but the error it points to is not guaranteed to be populated
// until after the iterator has been fully consumed.
func Err2[T, U any](f func(yield func(T, U) bool) error) (iter.Seq2[T, U], *error) {
	var err error
	result := func(yield func(T, U) bool) {
		err = f(yield)
	}
	return result, &err
}
