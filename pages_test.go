package seqs

import (
	"reflect"
	"slices"
	"testing"
)

func TestPage(t *testing.T) {
	var (
		ints    = Ints(1, 1)
		first10 = Limit(ints, 10)
		pages   = Pages(first10, 3)
		got     = slices.Collect(pages)
		want    = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}}
	)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFromPages(t *testing.T) {
	var (
		pages   = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}}
		pagenum = 0
		seq     = FromPages(func() ([]int, bool) {
			if pagenum >= len(pages) {
				return nil, false
			}
			result := pages[pagenum]
			pagenum++
			return result, pagenum <= len(pages)
		})
		got  = slices.Collect(seq)
		want = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
