package seqs

import (
	"slices"
	"testing"
)

func TestLimit(t *testing.T) {
	t.Run("long", func(t *testing.T) {
		var (
			ints  = Ints(0, 1)
			first = Limit(ints, 10)
			got   = slices.Collect(first)
			want  = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		)
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("short", func(t *testing.T) {
		var (
			ints  = slices.Values([]int{1, 2, 3})
			first = Limit(ints, 10)
			got   = slices.Collect(first)
			want  = []int{1, 2, 3}
		)
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestLimit2(t *testing.T) {
	t.Run("long", func(t *testing.T) {
		var (
			left   = Ints(0, 1)
			right  = Ints(100, 1)
			zipped = ZipVals(left, right)
			first  = Limit2(zipped, 10)
			pairs  = ToPairs(first)
			got    = slices.Collect(pairs)
			want   = []Pair[int, int]{
				{X: 0, Y: 100},
				{X: 1, Y: 101},
				{X: 2, Y: 102},
				{X: 3, Y: 103},
				{X: 4, Y: 104},
				{X: 5, Y: 105},
				{X: 6, Y: 106},
				{X: 7, Y: 107},
				{X: 8, Y: 108},
				{X: 9, Y: 109},
			}
		)
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("short", func(t *testing.T) {
		var (
			left   = slices.Values([]int{1, 2, 3})
			right  = slices.Values([]int{101, 102, 103})
			zipped = ZipVals(left, right)
			first  = Limit2(zipped, 10)
			pairs  = ToPairs(first)
			got    = slices.Collect(pairs)
			want   = []Pair[int, int]{
				{X: 1, Y: 101},
				{X: 2, Y: 102},
				{X: 3, Y: 103},
			}
		)
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestLast(t *testing.T) {
	var (
		ints     = Ints(0, 1)
		first100 = Limit(ints, 100)
		got      = LastN(first100, 7)
		want     = []int{93, 94, 95, 96, 97, 98, 99}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSkipN(t *testing.T) {
	t.Run("long", func(t *testing.T) {
		var (
			ints      = Ints(0, 1)
			skipped20 = SkipN(ints, 20)
			twenties  = Limit(skipped20, 10)
			got       = slices.Collect(twenties)
			want      = []int{20, 21, 22, 23, 24, 25, 26, 27, 28, 29}
		)
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("short", func(t *testing.T) {
		var (
			ints      = slices.Values([]int{1, 2, 3})
			skipped20 = SkipN(ints, 20)
			twenties  = Limit(skipped20, 10)
			got       = slices.Collect(twenties)
		)
		if len(got) > 0 {
			t.Errorf("got %v, want []", got)
		}
	})
}

func TestSkipN2(t *testing.T) {
	t.Run("long", func(t *testing.T) {
		var (
			left    = Ints(0, 1)
			right   = Ints(100, 1)
			zipped  = ZipVals(left, right)
			first   = Limit2(zipped, 20)
			skipped = SkipN2(first, 10)
			pairs   = ToPairs(skipped)
			got     = slices.Collect(pairs)
			want    = []Pair[int, int]{
				{X: 10, Y: 110},
				{X: 11, Y: 111},
				{X: 12, Y: 112},
				{X: 13, Y: 113},
				{X: 14, Y: 114},
				{X: 15, Y: 115},
				{X: 16, Y: 116},
				{X: 17, Y: 117},
				{X: 18, Y: 118},
				{X: 19, Y: 119},
			}
		)
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("short", func(t *testing.T) {
		var (
			left    = slices.Values([]int{1, 2, 3})
			right   = slices.Values([]int{101, 102, 103})
			zipped  = ZipVals(left, right)
			first   = Limit2(zipped, 20)
			skipped = SkipN2(first, 10)
			pairs   = ToPairs(skipped)
			got     = slices.Collect(pairs)
		)
		if len(got) > 0 {
			t.Errorf("got %v, want []", got)
		}
	})
}
