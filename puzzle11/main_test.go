package main

import "testing"

func TestPartOneExample(t *testing.T) {
	expected := 10605
	got := partOne(exampleMonkies())
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestPartTwoExample(t *testing.T) {
	expected := 2713310158
	got := partTwo(exampleMonkies())
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func exampleMonkies() []*monkey {
	return []*monkey{
		&monkey{
			items:         []int{79, 98},
			operation:     operation{op: opMultiply, operand: 19},
			testDivisible: 23,
			trueMonkey:    2,
			falseMonkey:   3,
		},
		&monkey{
			items:         []int{54, 65, 75, 74},
			operation:     operation{op: opAdd, operand: 6},
			testDivisible: 19,
			trueMonkey:    2,
			falseMonkey:   0,
		},
		&monkey{
			items:         []int{79, 60, 97},
			operation:     operation{op: opSquare},
			testDivisible: 13,
			trueMonkey:    1,
			falseMonkey:   3,
		},
		&monkey{
			items:         []int{74},
			operation:     operation{op: opAdd, operand: 3},
			testDivisible: 17,
			trueMonkey:    0,
			falseMonkey:   1,
		},
	}
}
