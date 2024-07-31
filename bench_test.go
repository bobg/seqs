package seqs

import "testing"

func BenchmarkIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var (
			ints      = Ints(1, 1)
			first1000 = FirstN(ints, 1000)
			odds      = Filter(first1000, func(x int) bool { return x%2 == 1 })
			squares   = Map(odds, func(x int) int { return x * x })
			sums      = Accum(squares, func(x, y int) int { return x + y })
		)
		sum := LastN(sums, 1)
		if len(sum) != 1 {
			b.Fatalf("len(sum) = %d, want 1", len(sum))
		}
	}
}
