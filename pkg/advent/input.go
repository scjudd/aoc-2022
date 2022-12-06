package advent

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/scjudd/aoc-2022/pkg/advent/internal/cache"
)

// GetInput will fetch the input for a puzzle and cache it such that it is only
// ever downloaded once.
func GetInput(session string, year, day int) (io.ReadCloser, error) {
	input, err := cache.GetInput(year, day)
	if err == nil {
		return input, nil
	} else if !cache.IsCacheMiss(err) {
		return nil, fmt.Errorf("error checking puzzle input cache: %w", err)
	}

	input, err = getLiveInput(session, year, day)
	if err != nil {
		return nil, fmt.Errorf("error getting live puzzle input: %w", err)
	}
	defer input.Close()

	copied := new(bytes.Buffer)
	cacheInput := io.TeeReader(input, copied)

	err = cache.SaveInput(year, day, cacheInput)
	if err != nil {
		return nil, fmt.Errorf("error updating puzzle input cache: %w", err)
	}

	return io.NopCloser(copied), nil
}

// MustGetInput is the same as GetInput, except it will panic on any errors.
func MustGetInput(session string, year, day int) io.ReadCloser {
	input, err := GetInput(session, year, day)
	if err != nil {
		panic(err)
	}
	return input
}

func getLiveInput(session string, year, day int) (io.ReadCloser, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Cookie", fmt.Sprintf("session=%s", session))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing HTTP request: %w", err)
	}
	return resp.Body, nil
}
