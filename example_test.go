package seqs_test

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/bobg/seqs"
)

func ExampleAccum() {
	var (
		ints   = seqs.Ints(1, 1)     // All integers starting at 1
		first5 = seqs.Limit(ints, 5) // First 5 integers
		sums   = seqs.Accum(first5, 0, func(a, b int) int { return a + b })
	)
	for val := range sums {
		fmt.Println(val)
	}
	// Output:
	// 1
	// 3
	// 6
	// 10
	// 15
}

func ExampleAccumx() {
	intstrs := slices.Values([]string{"1", "2", "3", "4", "5"})
	ints, errptr := seqs.Accumx(intstrs, 0, func(a int, s string) (int, error) {
		val, err := strconv.Atoi(s)
		return a + val, err
	})
	for sum := range ints {
		fmt.Println(sum)
	}
	if *errptr != nil {
		panic(*errptr)
	}
	// Output:
	// 1
	// 3
	// 6
	// 10
	// 15
}

func ExampleBytes() {
	const s = "Hello世界"
	for b := range seqs.Bytes(s) {
		fmt.Println(b)
	}
	// Output:
	// 72
	// 101
	// 108
	// 108
	// 111
	// 228
	// 184
	// 150
	// 231
	// 149
	// 140
}

func ExampleCheckEmpty() {
	var (
		ints  = seqs.Ints(1, 1) // All integers starting at 1
		empty = seqs.Empty[int] // No integers at all
	)

	ints, isEmpty := seqs.CheckEmpty(ints) // Reassign ints to get an iterator that preserves the first value of the original ints
	fmt.Printf("ints is empty: %t\n", isEmpty)

	empty, isEmpty = seqs.CheckEmpty(empty) // Reassign empty to get an iterator that preserves the first value of the original empty
	fmt.Printf("empty is empty: %t\n", isEmpty)

	_, _ = ints, empty // (silence linter "declared but not used" error)

	// Output:
	// ints is empty: false
	// empty is empty: true
}

func ExampleDup() {
	var (
		ints       = seqs.Ints(1, 1)      // All integers starting at 1
		first10    = seqs.Limit(ints, 10) // First 10 integers
		dups       = seqs.Dup(first10, 2) // Two copies of the first10 iterator
		evens      = seqs.Filter(dups[0], func(val int) bool { return val%2 == 0 })
		odds       = seqs.Filter(dups[1], func(val int) bool { return val%2 == 1 })
		evensSlice = slices.Collect(evens)
		oddsSlice  = slices.Collect(odds)
	)
	fmt.Println(evensSlice)
	fmt.Println(oddsSlice)
	// Output:
	// [2 4 6 8 10]
	// [1 3 5 7 9]
}

func ExampleEnumerate() {
	var (
		strs = slices.Values([]string{"a", "b", "c"})
		seq  = seqs.Enumerate(strs)
	)
	for i, s := range seq {
		fmt.Println(i, s)
	}
	// Output:
	// 0 a
	// 1 b
	// 2 c
}

func ExampleGo() {
	it, errptr := seqs.Go(func(ch chan<- int) error {
		ch <- 1
		ch <- 2
		ch <- 3
		return nil
	})
	slice := slices.Collect(it)
	if *errptr != nil {
		panic(*errptr)
	}
	fmt.Println(slice)
	// Output:
	// [1 2 3]
}

func ExampleLeft() {
	var (
		pairs = []seqs.Pair[int, string]{{X: 1, Y: "a"}, {X: 2, Y: "b"}, {X: 3, Y: "c"}}
		seq   = seqs.FromPairs(slices.Values(pairs))
		nums  = seqs.Left(seq)
	)
	fmt.Println(slices.Collect(nums))
	// Output:
	// [1 2 3]
}

func ExamplePages() {
	var (
		ints    = seqs.Ints(1, 1)      // All integers starting at 1
		first10 = seqs.Limit(ints, 10) // First 10 integers
	)
	pages := seqs.Pages(first10, 3)
	for page := range pages {
		fmt.Println(page)
	}
	// Output:
	// [1 2 3]
	// [4 5 6]
	// [7 8 9]
	// [10]
}

func ExamplePeek() {
	var (
		ints   = seqs.Ints(1, 1)     // All integers starting at 1
		first3 = seqs.Limit(ints, 3) // First three integers
	)

	peeked, ok, first3 := seqs.Peek(first3)
	if ok {
		all3 := slices.Collect(first3)
		fmt.Printf("Peeked value is %d, full sequence is %v\n", peeked, all3)
	} else {
		fmt.Println("Sequence is empty")
	}
	// Output:
	// Peeked value is 1, full sequence is [1 2 3]
}

func ExampleRight() {
	var (
		pairs = []seqs.Pair[int, string]{{X: 1, Y: "a"}, {X: 2, Y: "b"}, {X: 3, Y: "c"}}
		seq   = seqs.FromPairs(slices.Values(pairs))
		strs  = seqs.Right(seq)
	)
	fmt.Println(slices.Collect(strs))
	// Output:
	// [a b c]
}

func ExampleRunes() {
	const s = "Hello世界"
	for r := range seqs.Runes(s) {
		fmt.Println(string(r))
	}
	// Output:
	// H
	// e
	// l
	// l
	// o
	// 世
	// 界
}

func ExampleString() {
	const s = "Hello世界"
	for i, r := range seqs.String(s) {
		fmt.Printf("%d %c\n", i, r)
	}
	// Output:
	// 0 H
	// 1 e
	// 2 l
	// 3 l
	// 4 o
	// 5 世
	// 8 界
}

func ExampleUniq() {
	var (
		nums = slices.Values([]int{1, 1, 2, 2, 2, 3})
		uniq = seqs.Uniq(nums)
	)
	fmt.Println(slices.Collect(uniq))
	// Output:
	// [1 2 3]
}

func ExampleZipVals() {
	var (
		letters = slices.Values([]string{"a", "b", "c", "d"})
		nums    = slices.Values([]int{1, 2, 3})
		pairs   = seqs.ZipVals(letters, nums)
	)
	for x, y := range pairs {
		fmt.Println(x, y)
	}
	// Output:
	// a 1
	// b 2
	// c 3
	// d 0
}
