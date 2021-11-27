package tuple_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/tuple"
)

var (
	pair   = tuple.Pair[string, float64]{"abc", 3.14}
	triple = tuple.Triple[tuple.Pair[string, float64], int, []byte]{pair, 1, []byte{1, 2, 3}}

	_ tuple.Couple[string, float64] = tuple.Couple[string, float64](pair)
)

func TestPair(t *testing.T) {
	if pair.V1 != "abc" {
		t.Fatalf("V1 = %q, wants \"abc\"", pair.V1)
	}
	if pair.V2 != 3.14 {
		t.Fatalf("V2 = %v, wants 3.14", pair.V2)
	}
}

func TestTriple(t *testing.T) {
	if triple.V1 != pair {
		t.Fatalf("V1 = %#v, wants %#v", triple.V1, pair)
	}
	if triple.V2 != 1 {
		t.Fatalf("V2 = %v, wants 1", triple.V2)
	}
	if !reflect.DeepEqual(triple.V3, []byte{1, 2, 3}) {
		t.Fatalf("V3 = %v, wants %v", triple.V3, []byte{1, 2, 3})
	}
}
