package seqs

import (
	"bytes"
	"encoding/json/jsontext"
	"encoding/json/v2"
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestMarshalEncode(t *testing.T) {
	t.Run("non-empty", func(t *testing.T) {
		var (
			buf = new(bytes.Buffer)
			enc = jsontext.NewEncoder(buf)
		)

		seq := slices.Values([]int{1, 2, 3})
		if err := JSONMarshalEncode(enc, seq); err != nil {
			t.Fatal(err)
		}

		const want = "[1,2,3]"
		got := strings.TrimSpace(buf.String())
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("empty without option", func(t *testing.T) {
		var (
			buf = new(bytes.Buffer)
			enc = jsontext.NewEncoder(buf)
		)

		if err := JSONMarshalEncode(enc, Empty[int]); err != nil {
			t.Fatal(err)
		}

		const want = "[]"
		got := strings.TrimSpace(buf.String())
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("empty with FormatNilSliceAsNull true", func(t *testing.T) {
		var (
			buf = new(bytes.Buffer)
			enc = jsontext.NewEncoder(buf)
		)

		if err := JSONMarshalEncode(enc, Empty[int], json.FormatNilSliceAsNull(true)); err != nil {
			t.Fatal(err)
		}

		const want = "null"
		got := strings.TrimSpace(buf.String())
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("empty with FormatNilSliceAsNull false", func(t *testing.T) {
		var (
			buf = new(bytes.Buffer)
			enc = jsontext.NewEncoder(buf)
		)

		if err := JSONMarshalEncode(enc, Empty[int], json.FormatNilSliceAsNull(false)); err != nil {
			t.Fatal(err)
		}

		const want = "[]"
		got := strings.TrimSpace(buf.String())
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("empty with FormatNilSliceAsNull true in encoder options", func(t *testing.T) {
		var (
			buf = new(bytes.Buffer)
			enc = jsontext.NewEncoder(buf, json.FormatNilSliceAsNull(true))
		)

		if err := JSONMarshalEncode(enc, Empty[int]); err != nil {
			t.Fatal(err)
		}

		const want = "null"
		got := strings.TrimSpace(buf.String())
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("empty with FormatNilSliceAsNull true in encoder options overriden by call option", func(t *testing.T) {
		var (
			buf = new(bytes.Buffer)
			enc = jsontext.NewEncoder(buf, json.FormatNilSliceAsNull(true))
		)

		if err := JSONMarshalEncode(enc, Empty[int], json.FormatNilSliceAsNull(false)); err != nil {
			t.Fatal(err)
		}

		const want = "[]"
		got := strings.TrimSpace(buf.String())
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("error", func(t *testing.T) {
		var (
			buf = new(bytes.Buffer)
			enc = jsontext.NewEncoder(buf)
		)

		badSeq := func(yield func(errTypeVal) bool) {
			yield(errTypeVal{})
		}

		if err := JSONMarshalEncode(enc, badSeq); err == nil {
			t.Error("got nil error, want error")
		}
	})
}

func TestJSONMarshalEncode2(t *testing.T) {
	t.Run("non-empty", func(t *testing.T) {
		var (
			buf = new(bytes.Buffer)
			enc = jsontext.NewEncoder(buf)
		)

		seq := FromPairs(slices.Values([]Pair[string, int]{{"a", 1}, {"b", 2}}))
		if err := JSONMarshalEncode2(enc, seq); err != nil {
			t.Fatal(err)
		}

		const want = `{"a":1,"b":2}`
		got := strings.TrimSpace(buf.String())
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("empty", func(t *testing.T) {
		var (
			buf = new(bytes.Buffer)
			enc = jsontext.NewEncoder(buf)
		)

		if err := JSONMarshalEncode2(enc, Empty2[string, int]); err != nil {
			t.Fatal(err)
		}

		const want = "null"
		got := strings.TrimSpace(buf.String())
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("error", func(t *testing.T) {
		var (
			buf = new(bytes.Buffer)
			enc = jsontext.NewEncoder(buf)
		)

		badSeq := func(yield func(string, errTypeVal) bool) {
			yield("foo", errTypeVal{})
		}

		if err := JSONMarshalEncode2(enc, badSeq); err == nil {
			t.Error("got nil error, want error")
		}
	})
}

type errTypeVal struct{}

func (errTypeVal) MarshalJSON() ([]byte, error) {
	return nil, errors.New("marshal error")
}

func TestUnmarshalDecode(t *testing.T) {
	t.Run("valid array", func(t *testing.T) {
		dec := jsontext.NewDecoder(strings.NewReader("[1, 2, 3]"))
		seq, errptr := JSONUnmarshalDecode[int](dec)
		got := slices.Collect(seq)
		if err := *errptr; err != nil {
			t.Fatal(err)
		}
		want := []int{1, 2, 3}
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("null", func(t *testing.T) {
		dec := jsontext.NewDecoder(strings.NewReader("null"))
		seq, errptr := JSONUnmarshalDecode[int](dec)
		got := slices.Collect(seq)
		if err := *errptr; err != nil {
			t.Fatal(err)
		}
		if len(got) != 0 {
			t.Errorf("got %v, want empty", got)
		}
	})

	t.Run("invalid initial token kind", func(t *testing.T) {
		dec := jsontext.NewDecoder(strings.NewReader(`"not an array"`))
		_, errptr := JSONUnmarshalDecode[int](dec)
		err := *errptr
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if _, ok := errors.AsType[*json.SemanticError](err); !ok {
			t.Errorf("expected SemanticError, got %T: %v", err, err)
		}
	})

	t.Run("read initial token error", func(t *testing.T) {
		dec := jsontext.NewDecoder(strings.NewReader(""))
		_, errptr := JSONUnmarshalDecode[int](dec)
		if err := *errptr; err == nil {
			t.Fatal("expected error reading initial token, got nil")
		}
	})

	t.Run("element decode error", func(t *testing.T) {
		dec := jsontext.NewDecoder(strings.NewReader(`[1, "abc", 3]`))
		seq, errptr := JSONUnmarshalDecode[int](dec)
		got := slices.Collect(seq)
		if err := *errptr; err == nil {
			t.Fatal("expected error during element decode, got nil")
		}
		want := []int{1}
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("early stop", func(t *testing.T) {
		dec := jsontext.NewDecoder(strings.NewReader("[1, 2, 3, 4]"))
		seq, errptr := JSONUnmarshalDecode[int](dec)

		var got []int
		for val := range seq {
			got = append(got, val)
			if len(got) == 2 {
				break
			}
		}

		if err := *errptr; err != nil {
			t.Fatal(err)
		}

		want := []int{1, 2}
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestJSONUnmarshalDecode2(t *testing.T) {
	t.Run("valid object", func(t *testing.T) {
		dec := jsontext.NewDecoder(strings.NewReader(`{"a": 1, "b": 2}`))
		seq, errptr := JSONUnmarshalDecode2[int](dec)
		got := slices.Collect(ToPairs(seq))
		if err := *errptr; err != nil {
			t.Fatal(err)
		}
		want := []Pair[string, int]{{"a", 1}, {"b", 2}}
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("null", func(t *testing.T) {
		dec := jsontext.NewDecoder(strings.NewReader("null"))
		seq, errptr := JSONUnmarshalDecode2[int](dec)
		got := slices.Collect(ToPairs(seq))
		if err := *errptr; err != nil {
			t.Fatal(err)
		}
		if len(got) != 0 {
			t.Errorf("got %v, want empty", got)
		}
	})

	t.Run("invalid initial token kind", func(t *testing.T) {
		dec := jsontext.NewDecoder(strings.NewReader(`"not an object"`))
		_, errptr := JSONUnmarshalDecode2[int](dec)
		err := *errptr
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if _, ok := errors.AsType[*json.SemanticError](err); !ok {
			t.Errorf("expected SemanticError, got %T: %v", err, err)
		}
	})

	t.Run("read initial token error", func(t *testing.T) {
		dec := jsontext.NewDecoder(strings.NewReader(""))
		_, errptr := JSONUnmarshalDecode2[int](dec)
		if err := *errptr; err == nil {
			t.Fatal("expected error reading initial token, got nil")
		}
	})

	t.Run("non-string key error", func(t *testing.T) {
		dec := jsontext.NewDecoder(strings.NewReader(`{"a": 1, 123: 2}`))
		seq, errptr := JSONUnmarshalDecode2[int](dec)
		got := slices.Collect(ToPairs(seq))
		if err := *errptr; err == nil {
			t.Fatal("expected error for non-string key, got nil")
		}
		want := []Pair[string, int]{{"a", 1}}
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("element decode error", func(t *testing.T) {
		dec := jsontext.NewDecoder(strings.NewReader(`{"a": 1, "b": "abc", "c": 3}`))
		seq, errptr := JSONUnmarshalDecode2[int](dec)
		got := slices.Collect(ToPairs(seq))
		if err := *errptr; err == nil {
			t.Fatal("expected error during element decode, got nil")
		}
		want := []Pair[string, int]{{"a", 1}}
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("early stop", func(t *testing.T) {
		dec := jsontext.NewDecoder(strings.NewReader(`{"a": 1, "b": 2, "c": 3, "d": 4}`))
		seq, errptr := JSONUnmarshalDecode2[int](dec)

		var got []Pair[string, int]
		for k, v := range seq {
			got = append(got, Pair[string, int]{X: k, Y: v})
			if len(got) == 2 {
				break
			}
		}

		if err := *errptr; err != nil {
			t.Fatal(err)
		}

		want := []Pair[string, int]{{"a", 1}, {"b", 2}}
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestRoundTrip(t *testing.T) {
	var (
		input = []string{"foo", "bar", "baz"}
		inSeq = slices.Values(input)
		buf   = new(bytes.Buffer)
		enc   = jsontext.NewEncoder(buf)
	)

	if err := JSONMarshalEncode(enc, inSeq); err != nil {
		t.Fatalf("MarshalEncode error: %v", err)
	}

	dec := jsontext.NewDecoder(buf)
	outSeq, errptr := JSONUnmarshalDecode[string](dec)
	got := slices.Collect(outSeq)
	if err := *errptr; err != nil {
		t.Fatal(err)
	}

	if !slices.Equal(got, input) {
		t.Errorf("got %v, want %v", got, input)
	}
}

func TestRoundTrip2(t *testing.T) {
	var (
		input = []Pair[string, string]{{"foo", "a"}, {"bar", "b"}}
		inSeq = FromPairs(slices.Values(input))
		buf   = new(bytes.Buffer)
		enc   = jsontext.NewEncoder(buf)
	)

	if err := JSONMarshalEncode2(enc, inSeq); err != nil {
		t.Fatalf("JSONMarshalEncode2 error: %v", err)
	}

	dec := jsontext.NewDecoder(buf)
	outSeq, errptr := JSONUnmarshalDecode2[string](dec)
	got := slices.Collect(ToPairs(outSeq))
	if err := *errptr; err != nil {
		t.Fatal(err)
	}

	if !slices.Equal(got, input) {
		t.Errorf("got %v, want %v", got, input)
	}
}
