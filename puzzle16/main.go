package main

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/scjudd/aoc-2022/pkg/advent"
)

type (
	valve    = string
	flowRate = int
	distance = int
	bitmask  = int16

	puzzleData map[valve]puzzleDataElement

	puzzleDataElement struct {
		flowRate flowRate
		peers    []valve
	}

	searchState struct {
		origin valve
		time   int
		opened bitmask
	}
)

func main() {
	year, day := 2022, 16
	session := advent.MustLoadSession()
	data := parseInput(advent.MustGetInput(session, year, day))
	advent.PrintResult(advent.Check(session, year, day, 1, partOne(data)))
	advent.PrintResult(advent.Check(session, year, day, 2, partTwo(data)))
}

func partOne(data puzzleData) int {
	return pressureSearch(mapDistances(data), mapFlowRates(data), mapBitmaskIndex(data), make(map[searchState]int), searchState{"AA", 30, 0})
}

func partTwo(data puzzleData) int {
	distances := mapDistances(data)
	flowRates := mapFlowRates(data)
	bitmaskIndex := mapBitmaskIndex(data)
	memo := make(map[searchState]int)

	mask := bitmask(1<<len(flowRates) - 1)
	var best int
	for i := bitmask(0); i < mask/2; i++ {
		score := pressureSearch(distances, flowRates, bitmaskIndex, memo, searchState{"AA", 26, i})
		score += pressureSearch(distances, flowRates, bitmaskIndex, memo, searchState{"AA", 26, mask ^ i})
		if score > best {
			best = score
		}
	}
	return best
}

func pressureSearch(distances map[valve]map[valve]distance, flowRates map[valve]flowRate, bitmaskIndex map[valve]bitmask, memo map[searchState]int, state searchState) int {
	if best, ok := memo[state]; ok {
		return best
	}

	var best int
	for peer, distance := range distances[state.origin] {
		bit := bitmaskIndex[peer]
		if state.opened&bit != 0 {
			continue
		}
		next := searchState{
			origin: peer,
			time:   state.time - distance - 1,
			opened: state.opened | bit,
		}
		if next.time <= 0 {
			continue
		}
		if score := next.time*flowRates[next.origin] + pressureSearch(distances, flowRates, bitmaskIndex, memo, next); score > best {
			best = score
		}
	}
	memo[state] = best
	return best
}

// Return the distance from every non-zero flowRate valve to every other
// non-zero flowRate valve
func mapDistances(data puzzleData) map[valve]map[valve]distance {
	type queueItem struct {
		valve    valve
		distance distance
	}

	distances := make(map[valve]map[valve]distance)
	for v := range data {
		valveDistances := make(map[valve]distance)
		visited := map[valve]bool{v: true}
		queue := []queueItem{{v, 0}}
		for len(queue) > 0 {
			item := queue[0]
			queue = queue[1:]
			for _, peer := range data[item.valve].peers {
				if visited[peer] {
					continue
				}
				visited[peer] = true
				queue = append(queue, queueItem{peer, item.distance + 1})
				if data[peer].flowRate == 0 {
					continue
				}
				valveDistances[peer] = item.distance + 1
			}
		}
		distances[v] = valveDistances
	}
	return distances
}

func mapFlowRates(data puzzleData) map[valve]flowRate {
	flowRates := make(map[valve]flowRate)
	for valve, spec := range data {
		if spec.flowRate == 0 {
			continue
		}
		flowRates[valve] = spec.flowRate
	}
	return flowRates
}

func mapBitmaskIndex(data puzzleData) map[valve]bitmask {
	var i bitmask = 1
	bitmaskIndex := make(map[valve]bitmask)
	for valve, spec := range data {
		if spec.flowRate == 0 {
			continue
		}
		bitmaskIndex[valve] = i
		i <<= 1
	}
	return bitmaskIndex
}

func parseInput(input io.Reader) puzzleData {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	re := regexp.MustCompile("Valve ([A-Z]+) has flow rate=([0-9]+); tunnels? leads? to valves? ([A-Z, ]+)")
	data := make(puzzleData)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		flowRate, _ := strconv.Atoi(matches[2])
		data[matches[1]] = puzzleDataElement{
			flowRate: flowRate,
			peers:    strings.Split(matches[3], ", "),
		}
	}
	return data
}
