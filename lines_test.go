package seqs

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWords(t *testing.T) {
	indep, err := os.Open("testdata/indep.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer indep.Close()

	words, errptr := Words(indep)
	firstWords := FirstN(words, 10)
	got := slices.Collect(firstWords)
	if err := *errptr; err != nil {
		t.Fatal(err)
	}
	want := []string{
		"In", "Congress,", "July", "4,", "1776", "The", "unanimous", "Declaration", "of", "the",
	}
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLines(t *testing.T) {
	indep, err := os.ReadFile("testdata/indep.txt")
	if err != nil {
		t.Fatal(err)
	}

	lines, errptr := Lines(bytes.NewReader(indep))
	got := new(bytes.Buffer)
	for line := range lines {
		fmt.Fprintln(got, line)
	}
	if err := *errptr; err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(string(indep), got.String()); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestLongLines(t *testing.T) {
	indep, err := os.ReadFile("testdata/indep.txt")
	if err != nil {
		t.Fatal(err)
	}

	lines, errptr := LongLines(bytes.NewReader(indep))
	got := new(bytes.Buffer)

	for r := range lines {
		line, err := io.ReadAll(r)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Fprintln(got, string(line))
	}
	if err := *errptr; err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(string(indep), got.String()); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
