package seqs

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"iter"
)

// JSONMarshalEncode encodes a sequence of values as a JSON array.
// If the sequence is empty and the json.FormatNilSliceAsNull option is present in opts,
// it encodes a JSON null value instead. Otherwise, an empty sequence is encoded as an empty JSON array.
func JSONMarshalEncode[T any](out *jsontext.Encoder, in iter.Seq[T], opts ...json.Options) error {
	next, peek, stop := Peeker(in)
	defer stop()

	if _, ok := peek(); !ok {
		formatNil, _ := json.GetOption(json.JoinOptions(opts...), json.FormatNilSliceAsNull)
		if formatNil {
			return out.WriteToken(jsontext.Null)
		}
		if err := out.WriteToken(jsontext.BeginArray); err != nil {
			return err
		}
		return out.WriteToken(jsontext.EndArray)
	}

	if err := out.WriteToken(jsontext.BeginArray); err != nil {
		return err
	}

	for {
		val, ok := next()
		if !ok {
			return out.WriteToken(jsontext.EndArray)
		}
		if err := json.MarshalEncode(out, val, opts...); err != nil {
			return err
		}
	}
}

// JSONMarshalEncode2 encodes a sequence of key-value pairs as a JSON object.
// If the sequence is empty, it encodes a JSON null value instead.
// The keys must be strings, as required by the JSON object format.
func JSONMarshalEncode2[K ~string, V any](out *jsontext.Encoder, in iter.Seq2[K, V], opts ...json.Options) error {
	next, peek, stop := Peeker2(in)
	defer stop()

	if _, _, ok := peek(); !ok {
		return out.WriteToken(jsontext.Null)
	}

	if err := out.WriteToken(jsontext.BeginObject); err != nil {
		return err
	}

	for {
		key, val, ok := next()
		if !ok {
			return out.WriteToken(jsontext.EndObject)
		}
		if err := out.WriteToken(jsontext.String(string(key))); err != nil {
			return err
		}
		if err := json.MarshalEncode(out, val, opts...); err != nil {
			return err
		}
	}
}

// JSONUnmarshalDecode decodes a JSON array or null value into a sequence of values.
// If the input is a JSON null value, it returns an empty sequence.
// Unmarshaling happens incrementally as the resulting sequence is consumed.
// The caller may check for errors encountered during unmarshaling
// by dereferencing the returned error pointer after consuming the sequence.
func JSONUnmarshalDecode[T any](in *jsontext.Decoder, opts ...json.Options) (iter.Seq[T], *error) {
	tok, err := in.ReadToken()
	if err != nil {
		return Empty[T], &err
	}

	switch tok.Kind() {
	case jsontext.KindNull:
		return Empty[T], &err

	case jsontext.KindBeginArray:
		// continue below

	default:
		err = &json.SemanticError{Err: fmt.Errorf("expected JSON array or null, got %q", tok.String())}
		return Empty[T], &err
	}

	var outerErr error

	f := func(yield func(T) bool) {
		for in.PeekKind() != jsontext.KindEndArray {
			var val T
			if err := json.UnmarshalDecode(in, &val, opts...); err != nil {
				outerErr = err
				return
			}
			if !yield(val) {
				return
			}
		}
		_, err := in.ReadToken()
		outerErr = err
	}

	return f, &outerErr
}

// JSONUnmarshalDecode2 decodes a JSON object or null value into a sequence of key-value pairs.
// If the input is a JSON null value, it returns an empty sequence.
// Unmarshaling happens incrementally as the resulting sequence is consumed.
// The caller may check for errors encountered during unmarshaling
// by dereferencing the returned error pointer after consuming the sequence.
func JSONUnmarshalDecode2[T any](in *jsontext.Decoder, opts ...json.Options) (iter.Seq2[string, T], *error) {
	tok, err := in.ReadToken()
	if err != nil {
		return Empty2[string, T], &err
	}

	switch tok.Kind() {
	case jsontext.KindNull:
		return Empty2[string, T], &err

	case jsontext.KindBeginObject:
		// continue below

	default:
		err = &json.SemanticError{Err: fmt.Errorf("expected JSON object or null, got %q", tok.String())}
		return Empty2[string, T], &err
	}

	var outerErr error

	f := func(yield func(string, T) bool) {
		for in.PeekKind() != jsontext.KindEndObject {
			keyTok, err := in.ReadToken()
			if err != nil {
				outerErr = err
				return
			}
			if keyTok.Kind() != jsontext.KindString {
				outerErr = &json.SemanticError{Err: fmt.Errorf("expected JSON string for object key, got %q", keyTok.String())}
				return
			}
			key := keyTok.String()

			var val T
			if err := json.UnmarshalDecode(in, &val, opts...); err != nil {
				outerErr = err
				return
			}
			if !yield(key, val) {
				return
			}
		}
		_, err := in.ReadToken()
		outerErr = err
	}

	return f, &outerErr
}
