package seqs

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

func TestEqual(t *testing.T) {
	cases := []struct {
		inp1, inp2 []int
		want       bool
	}{
		{
			want: true,
		},
		{
			inp1: []int{1},
			inp2: []int{1},
			want: true,
		},
		{
			inp1: []int{1},
			inp2: []int{2},
			want: false,
		},
		{
			inp1: []int{1, 2},
			inp2: []int{1, 2},
			want: true,
		},
		{
			inp1: []int{1, 2},
			inp2: []int{2, 1},
			want: false,
		},
		{
			inp1: []int{1, 2},
			inp2: []int{1, 2, 3},
			want: false,
		},
		{
			inp1: []int{1, 2, 3},
			inp2: []int{1, 2},
			want: false,
		},
		{
			inp1: []int{1, 2, 3},
			inp2: []int{1, 2, 3},
			want: true,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			var (
				inp1 = slices.Values(tc.inp1)
				inp2 = slices.Values(tc.inp2)
				got  = Equal(inp1, inp2)
			)
			if got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestEqual2(t *testing.T) {
	cases := []struct {
		inp1, inp2 []int
		want       bool
	}{{
		want: true,
	}, {
		inp1: []int{1},
		inp2: []int{1},
		want: true,
	}, {
		inp1: []int{1},
		inp2: []int{2},
		want: false,
	}, {
		inp1: []int{1, 2},
		inp2: []int{1, 2},
		want: true,
	}, {
		inp1: []int{1, 2},
		inp2: []int{2, 1},
		want: false,
	}, {
		inp1: []int{1, 2},
		inp2: []int{1, 2, 3},
		want: false,
	}, {
		inp1: []int{1, 2, 3},
		inp2: []int{1, 2},
		want: false,
	}, {
		inp1: []int{1, 2, 3},
		inp2: []int{1, 2, 3},
		want: true,
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			var (
				inp1 = slices.All(tc.inp1)
				inp2 = slices.All(tc.inp2)
				got  = Equal2(inp1, inp2)
			)
			if got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestEqualFunc(t *testing.T) {
	cases := []struct {
		inp1, inp2 []string
		want       bool
	}{{
		want: true,
	}, {
		inp1: []string{"a"},
		inp2: []string{"a"},
		want: true,
	}, {
		inp1: []string{"a"},
		inp2: []string{"A"},
		want: true,
	}, {
		inp1: []string{"a"},
		inp2: []string{"b"},
		want: false,
	}, {
		inp1: []string{"a", "b"},
		inp2: []string{"a", "b"},
		want: true,
	}, {
		inp1: []string{"a", "B"},
		inp2: []string{"A", "b"},
		want: true,
	}, {
		inp1: []string{"a", "b"},
		inp2: []string{"b", "a"},
		want: false,
	}, {
		inp1: []string{"a", "b"},
		inp2: []string{"a", "b", "c"},
		want: false,
	}, {
		inp1: []string{"a", "b", "c"},
		inp2: []string{"a", "b"},
		want: false,
	}, {
		inp1: []string{"a", "b", "c"},
		inp2: []string{"a", "b", "c"},
		want: true,
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			var (
				inp1 = slices.Values(tc.inp1)
				inp2 = slices.Values(tc.inp2)
				got  = EqualFunc(inp1, inp2, func(x, y string) bool { return strings.EqualFold(x, y) })
			)
			if got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestEqualFunc2(t *testing.T) {
	cases := []struct {
		inp1, inp2 []string
		want       bool
	}{{
		want: true,
	}, {
		inp1: []string{"a"},
		inp2: []string{"a"},
		want: true,
	}, {
		inp1: []string{"a"},
		inp2: []string{"A"},
		want: true,
	}, {
		inp1: []string{"a"},
		inp2: []string{"b"},
		want: false,
	}, {
		inp1: []string{"a", "b"},
		inp2: []string{"a", "b"},
		want: true,
	}, {
		inp1: []string{"a", "B"},
		inp2: []string{"A", "b"},
		want: true,
	}, {
		inp1: []string{"a", "b"},
		inp2: []string{"b", "a"},
		want: false,
	}, {
		inp1: []string{"a", "b"},
		inp2: []string{"a", "b", "c"},
		want: false,
	}, {
		inp1: []string{"a", "b", "c"},
		inp2: []string{"a", "b"},
		want: false,
	}, {
		inp1: []string{"a", "b", "c"},
		inp2: []string{"a", "b", "c"},
		want: true,
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			var (
				inp1 = slices.All(tc.inp1)
				inp2 = slices.All(tc.inp2)
				got  = EqualFunc2(inp1, inp2, func(idx1 int, x string, idx2 int, y string) bool {
					return idx1 == idx2 && strings.EqualFold(x, y)
				})
			)
			if got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}
