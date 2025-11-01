package seqs

import (
	"encoding/json/jsontext"
	"errors"
	"fmt"
	"maps"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

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
		if n >= len(expectJSON) {
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
			t.Errorf("for pointer %q, got value %v (%T), want %v (%T)", pointer, val, val, wantVal, wantVal)
		}

		t.Logf("%q: %v\n", pointer, val)

		n++
	}

	if err := errors.Join(*errptr1, *errptr2); err != nil {
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
	"/0", true,
}, {
	"/1", false,
}, {
	"/2", Null{},
}, {
	"/3", map[string]any{},
}, {
	"/4", []any(nil),
}, {
	"", []any{true, false, Null{}, map[string]any{}, []any(nil)},
}, {
	"", "Remaining samples courtesy of Adobe: https://opensource.adobe.com/Spry/samples/data_region/JSONDataSetSample.html",
}, {
	"/0", int64(100),
}, {
	"/1", int64(500),
}, {
	"/2", int64(300),
}, {
	"/3", int64(200),
}, {
	"/4", int64(400),
}, {
	"", []any{int64(100), int64(500), int64(300), int64(200), int64(400)},
}, {
	"/0/color", "red",
}, {
	"/0/value", "#f00",
}, {
	"/0", map[string]any{"color": "red", "value": "#f00"},
}, {
	"/1/color", "green",
}, {
	"/1/value", "#0f0",
}, {
	"/1", map[string]any{"color": "green", "value": "#0f0"},
}, {
	"/2/color", "blue",
}, {
	"/2/value", "#00f",
}, {
	"/2", map[string]any{"color": "blue", "value": "#00f"},
}, {
	"/3/color", "cyan",
}, {
	"/3/value", "#0ff",
}, {
	"/3", map[string]any{"color": "cyan", "value": "#0ff"},
}, {
	"/4/color", "magenta",
}, {
	"/4/value", "#f0f",
}, {
	"/4", map[string]any{"color": "magenta", "value": "#f0f"},
}, {
	"/5/color", "yellow",
}, {
	"/5/value", "#ff0",
}, {
	"/5", map[string]any{"color": "yellow", "value": "#ff0"},
}, {
	"/6/color", "black",
}, {
	"/6/value", "#000",
}, {
	"/6", map[string]any{"color": "black", "value": "#000"},
}, {
	"", []any{
		map[string]any{"color": "red", "value": "#f00"},
		map[string]any{"color": "green", "value": "#0f0"},
		map[string]any{"color": "blue", "value": "#00f"},
		map[string]any{"color": "cyan", "value": "#0ff"},
		map[string]any{"color": "magenta", "value": "#f0f"},
		map[string]any{"color": "yellow", "value": "#ff0"},
		map[string]any{"color": "black", "value": "#000"},
	},
}, {
	"/color", "red",
}, {
	"/value", "#f00",
}, {
	"", map[string]any{"color": "red", "value": "#f00"},
}, {
	"/id", "0001",
}, {
	"/type", "donut",
}, {
	"/name", "Cake",
}, {
	"/ppu", 0.55,
}, {
	"/batters/batter/0/id", "1001",
}, {
	"/batters/batter/0/type", "Regular",
}, {
	"/batters/batter/0", map[string]any{"id": "1001", "type": "Regular"},
}, {
	"/batters/batter/1/id", "1002",
}, {
	"/batters/batter/1/type", "Chocolate",
}, {
	"/batters/batter/1", map[string]any{"id": "1002", "type": "Chocolate"},
}, {
	"/batters/batter/2/id", "1003",
}, {
	"/batters/batter/2/type", "Blueberry",
}, {
	"/batters/batter/2", map[string]any{"id": "1003", "type": "Blueberry"},
}, {
	"/batters/batter/3/id", "1004",
}, {
	"/batters/batter/3/type", "Devil's Food",
}, {
	"/batters/batter/3", map[string]any{"id": "1004", "type": "Devil's Food"},
}, {
	"/batters/batter", []any{
		map[string]any{"id": "1001", "type": "Regular"},
		map[string]any{"id": "1002", "type": "Chocolate"},
		map[string]any{"id": "1003", "type": "Blueberry"},
		map[string]any{"id": "1004", "type": "Devil's Food"},
	},
}, {
	"/batters", map[string]any{
		"batter": []any{
			map[string]any{"id": "1001", "type": "Regular"},
			map[string]any{"id": "1002", "type": "Chocolate"},
			map[string]any{"id": "1003", "type": "Blueberry"},
			map[string]any{"id": "1004", "type": "Devil's Food"},
		},
	},
}, {
	"/topping/0/id", "5001",
}, {
	"/topping/0/type", "None",
}, {
	"/topping/0", map[string]any{"id": "5001", "type": "None"},
}, {
	"/topping/1/id", "5002",
}, {
	"/topping/1/type", "Glazed",
}, {
	"/topping/1", map[string]any{"id": "5002", "type": "Glazed"},
}, {
	"/topping/2/id", "5005",
}, {
	"/topping/2/type", "Sugar",
}, {
	"/topping/2", map[string]any{"id": "5005", "type": "Sugar"},
}, {
	"/topping/3/id", "5007",
}, {
	"/topping/3/type", "Powdered Sugar",
}, {
	"/topping/3", map[string]any{"id": "5007", "type": "Powdered Sugar"},
}, {
	"/topping/4/id", "5006",
}, {
	"/topping/4/type", "Chocolate with Sprinkles",
}, {
	"/topping/4", map[string]any{"id": "5006", "type": "Chocolate with Sprinkles"},
}, {
	"/topping/5/id", "5003",
}, {
	"/topping/5/type", "Chocolate",
}, {
	"/topping/5", map[string]any{"id": "5003", "type": "Chocolate"},
}, {
	"/topping/6/id", "5004",
}, {
	"/topping/6/type", "Maple",
}, {
	"/topping/6", map[string]any{"id": "5004", "type": "Maple"},
}, {
	"/topping", []any{
		map[string]any{"id": "5001", "type": "None"},
		map[string]any{"id": "5002", "type": "Glazed"},
		map[string]any{"id": "5005", "type": "Sugar"},
		map[string]any{"id": "5007", "type": "Powdered Sugar"},
		map[string]any{"id": "5006", "type": "Chocolate with Sprinkles"},
		map[string]any{"id": "5003", "type": "Chocolate"},
		map[string]any{"id": "5004", "type": "Maple"},
	},
}, {
	"", map[string]any{
		"id":   "0001",
		"type": "donut",
		"name": "Cake",
		"ppu":  0.55,
		"batters": map[string]any{
			"batter": []any{
				map[string]any{"id": "1001", "type": "Regular"},
				map[string]any{"id": "1002", "type": "Chocolate"},
				map[string]any{"id": "1003", "type": "Blueberry"},
				map[string]any{"id": "1004", "type": "Devil's Food"},
			},
		},
		"topping": []any{
			map[string]any{"id": "5001", "type": "None"},
			map[string]any{"id": "5002", "type": "Glazed"},
			map[string]any{"id": "5005", "type": "Sugar"},
			map[string]any{"id": "5007", "type": "Powdered Sugar"},
			map[string]any{"id": "5006", "type": "Chocolate with Sprinkles"},
			map[string]any{"id": "5003", "type": "Chocolate"},
			map[string]any{"id": "5004", "type": "Maple"},
		},
	},
}, {
	"/0/id", "0001",
}, {
	"/0/type", "donut",
}, {
	"/0/name", "Cake",
}, {
	"/0/ppu", 0.55,
}, {
	"/0/batters/batter/0/id", "1001",
}, {
	"/0/batters/batter/0/type", "Regular",
}, {
	"/0/batters/batter/0", map[string]any{"id": "1001", "type": "Regular"},
}, {
	"/0/batters/batter/1/id", "1002",
}, {
	"/0/batters/batter/1/type", "Chocolate",
}, {
	"/0/batters/batter/1", map[string]any{"id": "1002", "type": "Chocolate"},
}, {
	"/0/batters/batter/2/id", "1003",
}, {
	"/0/batters/batter/2/type", "Blueberry",
}, {
	"/0/batters/batter/2", map[string]any{"id": "1003", "type": "Blueberry"},
}, {
	"/0/batters/batter/3/id", "1004",
}, {
	"/0/batters/batter/3/type", "Devil's Food",
}, {
	"/0/batters/batter/3", map[string]any{"id": "1004", "type": "Devil's Food"},
}, {
	"/0/batters/batter", []any{
		map[string]any{"id": "1001", "type": "Regular"},
		map[string]any{"id": "1002", "type": "Chocolate"},
		map[string]any{"id": "1003", "type": "Blueberry"},
		map[string]any{"id": "1004", "type": "Devil's Food"},
	},
}, {
	"/0/batters", map[string]any{
		"batter": []any{
			map[string]any{"id": "1001", "type": "Regular"},
			map[string]any{"id": "1002", "type": "Chocolate"},
			map[string]any{"id": "1003", "type": "Blueberry"},
			map[string]any{"id": "1004", "type": "Devil's Food"},
		},
	},
}, {
	"/0/topping/0/id", "5001",
}, {
	"/0/topping/0/type", "None",
}, {
	"/0/topping/0", map[string]any{"id": "5001", "type": "None"},
}, {
	"/0/topping/1/id", "5002",
}, {
	"/0/topping/1/type", "Glazed",
}, {
	"/0/topping/1", map[string]any{"id": "5002", "type": "Glazed"},
}, {
	"/0/topping/2/id", "5005",
}, {
	"/0/topping/2/type", "Sugar",
}, {
	"/0/topping/2", map[string]any{"id": "5005", "type": "Sugar"},
}, {
	"/0/topping/3/id", "5007",
}, {
	"/0/topping/3/type", "Powdered Sugar",
}, {
	"/0/topping/3", map[string]any{"id": "5007", "type": "Powdered Sugar"},
}, {
	"/0/topping/4/id", "5006",
}, {
	"/0/topping/4/type", "Chocolate with Sprinkles",
}, {
	"/0/topping/4", map[string]any{"id": "5006", "type": "Chocolate with Sprinkles"},
}, {
	"/0/topping/5/id", "5003",
}, {
	"/0/topping/5/type", "Chocolate",
}, {
	"/0/topping/5", map[string]any{"id": "5003", "type": "Chocolate"},
}, {
	"/0/topping/6/id", "5004",
}, {
	"/0/topping/6/type", "Maple",
}, {
	"/0/topping/6", map[string]any{"id": "5004", "type": "Maple"},
}, {
	"/0/topping", []any{
		map[string]any{"id": "5001", "type": "None"},
		map[string]any{"id": "5002", "type": "Glazed"},
		map[string]any{"id": "5005", "type": "Sugar"},
		map[string]any{"id": "5007", "type": "Powdered Sugar"},
		map[string]any{"id": "5006", "type": "Chocolate with Sprinkles"},
		map[string]any{"id": "5003", "type": "Chocolate"},
		map[string]any{"id": "5004", "type": "Maple"},
	},
}, {
	"/0", map[string]any{
		"batters": map[string]any{
			"batter": []any{
				map[string]any{
					"id":   "1001",
					"type": "Regular",
				},
				map[string]any{
					"id":   "1002",
					"type": "Chocolate",
				},
				map[string]any{
					"id":   "1003",
					"type": "Blueberry",
				},
				map[string]any{
					"id":   "1004",
					"type": "Devil's Food",
				},
			},
		},
		"id":   "0001",
		"name": "Cake",
		"ppu":  0.55,
		"topping": []any{
			map[string]any{
				"id":   "5001",
				"type": "None",
			},
			map[string]any{
				"id":   "5002",
				"type": "Glazed",
			},
			map[string]any{
				"id":   "5005",
				"type": "Sugar",
			},
			map[string]any{
				"id":   "5007",
				"type": "Powdered Sugar",
			},
			map[string]any{
				"id":   "5006",
				"type": "Chocolate with Sprinkles",
			},
			map[string]any{
				"id":   "5003",
				"type": "Chocolate",
			},
			map[string]any{
				"id":   "5004",
				"type": "Maple",
			},
		},
		"type": "donut",
	},
}, {
	"/1/id", "0002",
}, {
	"/1/type", "donut",
}, {
	"/1/name", "Raised",
}, {
	"/1/ppu", 0.55,
}, {
	"/1/batters/batter/0/id", "1001",
}, {
	"/1/batters/batter/0/type", "Regular",
}, {
	"/1/batters/batter/0", map[string]any{
		"id":   "1001",
		"type": "Regular",
	},
}, {
	"/1/batters/batter", []any{
		map[string]any{
			"id":   "1001",
			"type": "Regular",
		},
	},
}, {
	"/1/batters", map[string]any{
		"batter": []any{
			map[string]any{
				"id":   "1001",
				"type": "Regular",
			},
		},
	},
}, {
	"/1/topping/0/id", "5001",
}, {
	"/1/topping/0/type", "None",
}, {
	"/1/topping/0", map[string]any{
		"id":   "5001",
		"type": "None",
	},
}, {
	"/1/topping/1/id", "5002",
}, {
	"/1/topping/1/type", "Glazed",
}, {
	"/1/topping/1", map[string]any{
		"id":   "5002",
		"type": "Glazed",
	},
}, {
	"/1/topping/2/id", "5005",
}, {
	"/1/topping/2/type", "Sugar",
}, {
	"/1/topping/2", map[string]any{
		"id":   "5005",
		"type": "Sugar",
	},
}, {
	"/1/topping/3/id", "5003",
}, {
	"/1/topping/3/type", "Chocolate",
}, {
	"/1/topping/3", map[string]any{
		"id":   "5003",
		"type": "Chocolate",
	},
}, {
	"/1/topping/4/id", "5004",
}, {
	"/1/topping/4/type", "Maple",
}, {
	"/1/topping/4", map[string]any{
		"id":   "5004",
		"type": "Maple",
	},
}, {
	"/1/topping", []any{
		map[string]any{
			"id":   "5001",
			"type": "None",
		},
		map[string]any{
			"id":   "5002",
			"type": "Glazed",
		},
		map[string]any{
			"id":   "5005",
			"type": "Sugar",
		},
		map[string]any{
			"id":   "5003",
			"type": "Chocolate",
		},
		map[string]any{
			"id":   "5004",
			"type": "Maple",
		},
	},
}, {
	"/1", map[string]any{
		"batters": map[string]any{
			"batter": []any{
				map[string]any{
					"id":   "1001",
					"type": "Regular",
				},
			},
		},
		"id":   "0002",
		"name": "Raised",
		"ppu":  0.55,
		"topping": []any{
			map[string]any{
				"id":   "5001",
				"type": "None",
			},
			map[string]any{
				"id":   "5002",
				"type": "Glazed",
			},
			map[string]any{
				"id":   "5005",
				"type": "Sugar",
			},
			map[string]any{
				"id":   "5003",
				"type": "Chocolate",
			},
			map[string]any{
				"id":   "5004",
				"type": "Maple",
			},
		},
		"type": "donut",
	},
}, {
	"/2/id", "0003",
}, {
	"/2/type", "donut",
}, {
	"/2/name", "Old Fashioned",
}, {
	"/2/ppu", 0.55,
}, {
	"/2/batters/batter/0/id", "1001",
}, {
	"/2/batters/batter/0/type", "Regular",
}, {
	"/2/batters/batter/0", map[string]any{
		"id":   "1001",
		"type": "Regular",
	},
}, {
	"/2/batters/batter/1/id", "1002",
}, {
	"/2/batters/batter/1/type", "Chocolate",
}, {
	"/2/batters/batter/1", map[string]any{
		"id":   "1002",
		"type": "Chocolate",
	},
}, {
	"/2/batters/batter", []any{
		map[string]any{
			"id":   "1001",
			"type": "Regular",
		},
		map[string]any{
			"id":   "1002",
			"type": "Chocolate",
		},
	},
}, {
	"/2/batters", map[string]any{
		"batter": []any{
			map[string]any{
				"id":   "1001",
				"type": "Regular",
			},
			map[string]any{
				"id":   "1002",
				"type": "Chocolate",
			},
		},
	},
}, {
	"/2/topping/0/id", "5001",
}, {
	"/2/topping/0/type", "None",
}, {
	"/2/topping/0", map[string]any{
		"id":   "5001",
		"type": "None",
	},
}, {
	"/2/topping/1/id", "5002",
}, {
	"/2/topping/1/type", "Glazed",
}, {
	"/2/topping/1", map[string]any{
		"id":   "5002",
		"type": "Glazed",
	},
}, {
	"/2/topping/2/id", "5003",
}, {
	"/2/topping/2/type", "Chocolate",
}, {
	"/2/topping/2", map[string]any{
		"id":   "5003",
		"type": "Chocolate",
	},
}, {
	"/2/topping/3/id", "5004",
}, {
	"/2/topping/3/type", "Maple",
}, {
	"/2/topping/3", map[string]any{
		"id":   "5004",
		"type": "Maple",
	},
}, {
	"/2/topping", []any{
		map[string]any{
			"id":   "5001",
			"type": "None",
		},
		map[string]any{
			"id":   "5002",
			"type": "Glazed",
		},
		map[string]any{
			"id":   "5003",
			"type": "Chocolate",
		},
		map[string]any{
			"id":   "5004",
			"type": "Maple",
		},
	},
}, {
	"/2", map[string]any{
		"batters": map[string]any{
			"batter": []any{
				map[string]any{
					"id":   "1001",
					"type": "Regular",
				},
				map[string]any{
					"id":   "1002",
					"type": "Chocolate",
				},
			},
		},
		"id":   "0003",
		"name": "Old Fashioned",
		"ppu":  0.55,
		"topping": []any{
			map[string]any{
				"id":   "5001",
				"type": "None",
			},
			map[string]any{
				"id":   "5002",
				"type": "Glazed",
			},
			map[string]any{
				"id":   "5003",
				"type": "Chocolate",
			},
			map[string]any{
				"id":   "5004",
				"type": "Maple",
			},
		},
		"type": "donut",
	},
}, {
	"", []any{
		map[string]any{
			"batters": map[string]any{
				"batter": []any{
					map[string]any{
						"id":   "1001",
						"type": "Regular",
					},
					map[string]any{
						"id":   "1002",
						"type": "Chocolate",
					},
					map[string]any{
						"id":   "1003",
						"type": "Blueberry",
					},
					map[string]any{
						"id":   "1004",
						"type": "Devil's Food",
					},
				},
			},
			"id":   "0001",
			"name": "Cake",
			"ppu":  0.55,
			"topping": []any{
				map[string]any{
					"id":   "5001",
					"type": "None",
				},
				map[string]any{
					"id":   "5002",
					"type": "Glazed",
				},
				map[string]any{
					"id":   "5005",
					"type": "Sugar",
				},
				map[string]any{
					"id":   "5007",
					"type": "Powdered Sugar",
				},
				map[string]any{
					"id":   "5006",
					"type": "Chocolate with Sprinkles",
				},
				map[string]any{
					"id":   "5003",
					"type": "Chocolate",
				},
				map[string]any{
					"id":   "5004",
					"type": "Maple",
				},
			},
			"type": "donut",
		},
		map[string]any{
			"batters": map[string]any{
				"batter": []any{
					map[string]any{
						"id":   "1001",
						"type": "Regular",
					},
				},
			},
			"id":   "0002",
			"name": "Raised",
			"ppu":  0.55,
			"topping": []any{
				map[string]any{
					"id":   "5001",
					"type": "None",
				},
				map[string]any{
					"id":   "5002",
					"type": "Glazed",
				},
				map[string]any{
					"id":   "5005",
					"type": "Sugar",
				},
				map[string]any{
					"id":   "5003",
					"type": "Chocolate",
				},
				map[string]any{
					"id":   "5004",
					"type": "Maple",
				},
			},
			"type": "donut",
		},
		map[string]any{
			"batters": map[string]any{
				"batter": []any{
					map[string]any{
						"id":   "1001",
						"type": "Regular",
					},
					map[string]any{
						"id":   "1002",
						"type": "Chocolate",
					},
				},
			},
			"id":   "0003",
			"name": "Old Fashioned",
			"ppu":  0.55,
			"topping": []any{
				map[string]any{
					"id":   "5001",
					"type": "None",
				},
				map[string]any{
					"id":   "5002",
					"type": "Glazed",
				},
				map[string]any{
					"id":   "5003",
					"type": "Chocolate",
				},
				map[string]any{
					"id":   "5004",
					"type": "Maple",
				},
			},
			"type": "donut",
		},
	},
}}

func fmtpair(pointer jsontext.Pointer, val any) string {
	return fmt.Sprintf("  %q, %s", pointer, fmtval(val, 0))
}

func fmtval(val any, depth int) string {
	switch val := val.(type) {
	case string:
		return strconv.Quote(val)

	case []any:
		if len(val) == 0 {
			return "[]any{}"
		}

		var sb strings.Builder
		fmt.Fprint(&sb, "[]any{\n")
		for _, elt := range val {
			fmt.Fprintf(&sb, "  %s%s,\n", strings.Repeat("  ", depth+1), fmtval(elt, depth+1))
		}
		fmt.Fprintf(&sb, "  %s}", strings.Repeat("  ", depth))
		return sb.String()

	case map[string]any:
		if len(val) == 0 {
			return "map[string]any{}"
		}

		var sb strings.Builder
		fmt.Fprint(&sb, "map[string]any{\n")
		keys := slices.Collect(maps.Keys(val))
		sort.Strings(keys)
		for _, key := range keys {
			fmt.Fprintf(&sb, "  %s%q: %s,\n", strings.Repeat("  ", depth+1), key, fmtval(val[key], depth+1))
		}
		fmt.Fprintf(&sb, "  %s}", strings.Repeat("  ", depth))
		return sb.String()

	default:
		return fmt.Sprintf("%v", val)
	}
}
