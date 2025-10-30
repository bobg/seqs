package seqs

import (
	"encoding/json/jsontext"
	"errors"
	"fmt"
	"io"
	"iter"
	"math"
	"strconv"
)

// JSONTokens parses JSON tokens from r and returns them as an [iter.Seq].
// This sequence is suitable as input to [JSONValues].
//
// The caller may check for errors by dereferencing the returned error pointer,
// but only after the iterator is fully consumed.
func JSONTokens(r io.Reader, opts ...jsontext.Options) (iter.Seq[jsontext.Token], *error) {
	var (
		dec      = jsontext.NewDecoder(r, opts...)
		outerErr error
	)
	f := func(yield func(jsontext.Token) bool) {
		for {
			tok, err := dec.ReadToken()
			if errors.Is(err, io.EOF) {
				return
			}
			if err != nil {
				outerErr = err
				return
			}
			if !yield(tok) {
				return
			}
		}
	}
	return f, &outerErr
}

// JSONValues consumes a sequence of JSON tokens and produces a sequence of JSON values,
// each paired with the [jsontext.Pointer] that can locate it within its top-level object.
//
// Values are produced as they are encountered, in depth-first fashion,
// making this a "streaming" or "event-based" parser.
// For example, given a sequence of tokens representing this input:
//
//	{"hello": [1, 2], "world": [3, 4]}
//
// JSONValues will produce pointer/value pairs in this order:
//
//	"/hello/0"  1
//	"/hello/1"  2
//	"/hello"    [1, 2]
//	"/world/0"  3
//	"/world/1"  4
//	"/world"    [3, 4]
//	""          {"hello": [1, 2], "world": [3, 4]}
//
// Note that object keys are not considered values to be separately emitted.
//
// The input to this function may be supplied by a call to [JSONTokens].
//
// Value types in the resulting sequence are:
//
//   - []any for arrays
//   - map[string]any for objects
//   - strings for strings
//   - boolean for booleans
//   - any(nil) for null
//
// and, for numbers:
//
//   - int64, if it can represent the value without loss of precision; otherwise
//   - uint64, if that can; otherwise
//   - float64.
//
// The input may contain multiple top-level JSON value.
// If the input ends in the middle of a JSON value,
// JSONValues produces an [io.ErrUnexpectedEOF] error.
//
// The caller may check for errors by dereferencing the returned error pointer,
// but only after the iterator is fully consumed.
func JSONValues(tokens iter.Seq[jsontext.Token]) (iter.Seq2[jsontext.Pointer, any], *error) {
	var outerErr error

	f := func(yield func(jsontext.Pointer, any) bool) {
		var (
			stack      []*stackItem
			nextObjKey *string // nil means expecting a map key (when top-of-stack is a map)
		)

		for tok := range tokens {
			var (
				kind = tok.Kind()
				val  any
				str  string
			)

			switch kind {
			case 'n':
				// val remains nil

			case 'f':
				val = false

			case 't':
				val = true

			case '"':
				str = tok.String()
				val = str

			case '0':
				num, err := parseJSONNum(tok)
				if err != nil {
					outerErr = err
					return
				}
				val = num

			case '{':
				stack = append(stack, &stackItem{val: make(map[string]any)})
				continue

			case '}':
				if len(stack) == 0 {
					outerErr = fmt.Errorf("unexpected close brace: stack empty")
					return
				}
				top := stack[len(stack)-1]
				obj, ok := top.val.(map[string]any)
				if !ok {
					outerErr = fmt.Errorf("unexpected close brace in non-object")
					return
				}
				val = obj
				stack = stack[:len(stack)-1]

			case '[':
				stack = append(stack, &stackItem{val: []any(nil)})
				continue

			case ']':
				if len(stack) == 0 {
					outerErr = fmt.Errorf("unexpected close bracket: stack empty")
					return
				}
				top := stack[len(stack)-1]
				obj, ok := top.val.([]any)
				if !ok {
					outerErr = fmt.Errorf("unexpected close bracket in non-array")
					return
				}
				val = obj
				stack = stack[:len(stack)-1]

			default:
				outerErr = fmt.Errorf("unknown token kind '%v'", kind)
				return
			}

			if len(stack) > 0 {
				top := stack[len(stack)-1]
				switch topval := top.val.(type) {
				case map[string]any:
					if nextObjKey == nil {
						if kind != '"' {
							outerErr = fmt.Errorf("got %s token, want string", kind)
							return
						}
						nextObjKey = &str
						top.key = str
						continue
					}
					topval[*nextObjKey] = val
					nextObjKey = nil

				case []any:
					topval = append(topval, val)
					top.val = topval

				default:
					outerErr = fmt.Errorf("internal error: unexpected %T on the stack", top)
					return
				}
			}

			var pointer jsontext.Pointer
			for _, s := range stack {
				switch sval := s.val.(type) {
				case map[string]any:
					pointer = pointer.AppendToken(s.key)

				case []any:
					pointer = pointer.AppendToken(strconv.Itoa(len(sval) - 1))

				default:
					outerErr = fmt.Errorf("internal error: unexpected %T on stack", sval)
					return
				}
			}

			if !yield(pointer, val) {
				return
			}
		}

		if len(stack) > 0 {
			outerErr = io.ErrUnexpectedEOF
			return
		}
	}

	return f, &outerErr
}

type stackItem struct {
	val any    // []any for arrays, map[string]any for objs
	key string // when val is a map[string]any, this is the latest key seen
}

// Returns an int64 if possible, otherwise a uint64 if possible, otherwise a float64.
func parseJSONNum(tok jsontext.Token) (_ any, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = fmt.Errorf("getting float value of JSON token: %w", e)
			} else {
				err = fmt.Errorf("getting float value of JSON token: %v", r)
			}
		}
	}()

	f := tok.Float()
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return f, nil
	}

	if r := math.Round(f); r != f {
		return f, nil
	}

	if f >= math.MinInt && f <= math.MaxInt {
		return tok.Int(), nil
	}

	if f >= 0 && f <= math.MaxUint {
		return tok.Uint(), nil
	}

	return f, nil
}
