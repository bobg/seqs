package seqs_test

import (
	"fmt"
	"maps"
	"slices"
	"sort"
	"sync"

	"github.com/bobg/seqs"
)

func ExampleAccum() {
	var (
		ints   = seqs.Ints(1, 1)      // All integers starting at 1
		first5 = seqs.FirstN(ints, 5) // First 5 integers
		sums   = seqs.Accum(first5, func(a, b int) int { return a + b })
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

func ExampleDup() {
	var (
		ints       = seqs.Ints(1, 1)       // All integers starting at 1
		first10    = seqs.FirstN(ints, 10) // First 10 integers
		dups       = seqs.Dup(first10, 2)  // Two copies of the first10 iterator
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
		ints    = seqs.Ints(1, 1)       // All integers starting at 1
		first10 = seqs.FirstN(ints, 10) // First 10 integers
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

func ExamplePartition() {
	var (
		ints       = seqs.Ints(0, 1)                                           // All integers starting at 0
		first10    = seqs.FirstN(ints, 10)                                     // First 10 integers
		partitions = seqs.Partition(first10, func(n int) int { return n % 3 }) // 3 subsequences: integers mod 3

		wg sync.WaitGroup
		mu sync.Mutex // Protects m
		m  = make(map[int][]int)
	)

	// This loop illustrates how the subsequences produced by Partition should be consumed in parallel.
	for k, seq := range partitions {
		wg.Add(1)
		go func() {
			members := slices.Collect(seq)
			mu.Lock()
			m[k] = members
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	keys := slices.Collect(maps.Keys(m))
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println(k, m[k])
	}

	// Output:
	// 0 [0 3 6 9]
	// 1 [1 4 7]
	// 2 [2 5 8]
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

func ExampleZip() {
	var (
		letters = slices.Values([]string{"a", "b", "c", "d"})
		nums    = slices.Values([]int{1, 2, 3})
		pairs   = seqs.Zip(letters, nums)
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
