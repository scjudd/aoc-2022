package main

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPartOneExample(t *testing.T) {
	expected := 1651
	got := partOne(exampleInput())
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestPartTwoExample(t *testing.T) {
	expected := 1707
	got := partTwo(exampleInput())
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestParseInput(t *testing.T) {
	input := strings.NewReader(`Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`)
	expected := exampleInput()
	got := parseInput(input)
	if diff := cmp.Diff(expected, got, cmp.AllowUnexported(puzzleDataElement{})); diff != "" {
		t.Errorf("Input parsing failed, (-expected,+got)\n%s", diff)
	}
}

func exampleInput() puzzleData {
	return puzzleData{
		"AA": {flowRate: 0, peers: []string{"DD", "II", "BB"}},
		"BB": {flowRate: 13, peers: []string{"CC", "AA"}},
		"CC": {flowRate: 2, peers: []string{"DD", "BB"}},
		"DD": {flowRate: 20, peers: []string{"CC", "AA", "EE"}},
		"EE": {flowRate: 3, peers: []string{"FF", "DD"}},
		"FF": {flowRate: 0, peers: []string{"EE", "GG"}},
		"GG": {flowRate: 0, peers: []string{"FF", "HH"}},
		"HH": {flowRate: 22, peers: []string{"GG"}},
		"II": {flowRate: 0, peers: []string{"AA", "JJ"}},
		"JJ": {flowRate: 21, peers: []string{"II"}},
	}
}
