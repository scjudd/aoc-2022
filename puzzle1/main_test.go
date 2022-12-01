package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestPartOneExample(t *testing.T) {
	calorieList := [][]int{
		[]int{1000, 2000, 3000},
		[]int{4000},
		[]int{5000, 6000},
		[]int{7000, 8000, 9000},
		[]int{10000},
	}
	expected := 24000
	got := partOne(calorieList)
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestPartTwoExample(t *testing.T) {
	calorieList := [][]int{
		[]int{1000, 2000, 3000},
		[]int{4000},
		[]int{5000, 6000},
		[]int{7000, 8000, 9000},
		[]int{10000},
	}
	expected := 45000
	got := partTwo(calorieList)
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestParseInput(t *testing.T) {
	input := strings.NewReader("1\n2\n\n3\n\n4\n5\n6")
	expected := [][]int{[]int{1, 2}, []int{3}, []int{4, 5, 6}}
	got := parseInput(input)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %#v, got %#v", expected, got)
	}
}

func TestMaxList(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	expected := []int{6, 5, 4}
	got := []int{0, 0, 0}
	for _, num := range list {
		maxList(got, num)
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %#v, got %#v", expected, got)
	}
}
