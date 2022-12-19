package main

import "testing"

func TestPartOneExample(t *testing.T) {
	expected := _part_one_expected_
	got := partOne(exampleData())
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

// func TestPartTwoExample(t *testing.T) {
// 	expected := _part_two_expected_
// 	got := partTwo(exampleData())
// 	if expected != got {
// 		t.Errorf("Expected %d, got %d", expected, got)
// 	}
// }

// func TestInputParsing(t *testing.T) {
// 	t.Fatal("not implemented")
// }

func exampleData() []int {
	return []int{
		0,
	}
}
