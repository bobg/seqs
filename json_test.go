package seqs

import (
	"encoding/json/jsontext"
	"os"
	"reflect"
	"testing"
)

func TestJSONValues(t *testing.T) {
	inp, err := os.Open("testdata/input.json")
	if err != nil {
		t.Fatal(err)
	}
	defer inp.Close()

	toks, errptr1 := JSONTokens(inp)
	pairs, errptr2 := JSONValues(toks)

	var n int

	for pointer, val := range pairs {
		if n > len(expectJSON) {
			t.Fatalf(`not enough "expect" pairs after %d values`, n)
		}

		var (
			wantPointer = expectJSON[n].p
			wantVal     = expectJSON[n].v
		)

		if pointer != wantPointer {
			t.Errorf("got pointer %q, want %q", pointer, wantPointer)
		}
		if !reflect.DeepEqual(val, wantVal) {
			t.Errorf("got value %v (%T), want %v (%T)", val, val, wantVal, wantVal)
		}

		t.Logf("%q %v\n", pointer, val)

		n++
	}

	if err := *errptr2; err != nil {
		t.Fatal(err)
	}
	if err := *errptr1; err != nil {
		t.Fatal(err)
	}

	if n < len(expectJSON) {
		t.Fatalf(`extra "want" tuple(s) after %d values`, n)
	}
}

var expectJSON = []struct {
	p jsontext.Pointer
	v any
}{{
	"", int64(17),
}, {
	"/0", int64(1),
}, {
	"/1", int64(7),
}, {
	"/2", "xyz",
}, {
	"", []any{int64(1), int64(7), "xyz"},
}}
