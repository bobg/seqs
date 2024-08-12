# Seqs - functions for operating on and with Go iterable sequences

[![Go Reference](https://pkg.go.dev/badge/github.com/bobg/seqs.svg)](https://pkg.go.dev/github.com/bobg/seqs)
[![Go Report Card](https://goreportcard.com/badge/github.com/bobg/seqs)](https://goreportcard.com/report/github.com/bobg/seqs)
[![Tests](https://github.com/bobg/seqs/actions/workflows/go.yml/badge.svg)](https://github.com/bobg/seqs/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/bobg/seqs/badge.svg?branch=main)](https://coveralls.io/github/bobg/seqs?branch=main)

This is seqs,
a library for working with iterable sequences,
introduced via the “range over function” language feature
and the [iter](https://pkg.go.dev/iter) standard-library package
in Go 1.23.

This is an adaptation to Go sequences of [github.com/bobg/go-generics/v3/iter](https://pkg.go.dev/github.com/bobg/go-generics/v3/iter).
The central type in that package,
[iter.Of](https://pkg.go.dev/github.com/bobg/go-generics/v3/iter#Of),
was made obsolete with the introduction of Go sequences.
This library provides operations on sequences analogous to the ones on `iter.Of` in the old package.
It combines those with [the proposal for an x/exp/iter package](https://github.com/golang/go/issues/61898).
