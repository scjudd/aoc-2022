package main

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPartOneExample(t *testing.T) {
	expected := 13
	got := partOne(examplePackets())
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestPartTwoExample(t *testing.T) {
	expected := 140
	got := partTwo(examplePackets())
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestTokenize(t *testing.T) {
	s := "[1,[20,3],4]"
	expected := []string{"[", "1", "[", "20", "3", "]", "4", "]"}
	got := tokenize(s)
	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("tokenize returned unexpected result (-expected,+got)\n%s", diff)
	}
}

func TestParsePacket(t *testing.T) {
	s := "[1,[20,3],4]"
	expected := l(i(1), l(i(20), i(3)), i(4))
	got := parsePacket(s)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func examplePackets() []*packet {
	return []*packet{
		// [1,1,3,1,1]
		// [1,1,5,1,1]
		l(i(1), i(1), i(3), i(1), i(1)),
		l(i(1), i(1), i(5), i(1), i(1)),

		// [[1],[2,3,4]]
		// [[1],4]
		l(l(i(1)), l(i(2), i(3), i(4))),
		l(l(i(1)), i(4)),

		// [9]
		// [[8,7,6]]
		l(i(9)),
		l(l(i(8), i(7), i(6))),

		// [[4,4],4,4]
		// [[4,4],4,4,4]
		l(l(i(4), i(4)), i(4), i(4)),
		l(l(i(4), i(4)), i(4), i(4), i(4)),

		// [7,7,7,7]
		// [7,7,7]
		l(i(7), i(7), i(7), i(7)),
		l(i(7), i(7), i(7)),

		// []
		// [3]
		l(),
		l(i(3)),

		// [[[]]]
		// [[]]
		l(l(l())),
		l(l()),

		// [1,[2,[3,[4,[5,6,7]]]],8,9]
		// [1,[2,[3,[4,[5,6,0]]]],8,9]
		l(i(1), l(i(2), l(i(3), l(i(4), l(i(5), i(6), i(7))))), i(8), i(9)),
		l(i(1), l(i(2), l(i(3), l(i(4), l(i(5), i(6), i(0))))), i(8), i(9)),
	}
}
