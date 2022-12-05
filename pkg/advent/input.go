package advent

import (
	"bytes"
	"fmt"
	"github.com/scjudd/aoc-2022/pkg/advent/internal/cache"
	"io"
	"net/http"
)

func GetInput(s *State) (io.ReadCloser, error) {
	input, err := cache.GetInput(s.Year, s.Day)
	if err == nil {
		return input, nil
	} else if !cache.IsCacheMiss(err) {
		return nil, fmt.Errorf("error checking puzzle input cache: %w", err)
	}

	input, err = getLiveInput(s.Session, s.Year, s.Day)
	if err != nil {
		return nil, fmt.Errorf("error getting live puzzle input: %w", err)
	}
	defer input.Close()

	copied := new(bytes.Buffer)
	cacheInput := io.TeeReader(input, copied)

	err = cache.SaveInput(s.Year, s.Day, cacheInput)
	if err != nil {
		return nil, fmt.Errorf("error updating puzzle input cache: %w", err)
	}

	return io.NopCloser(copied), nil
}

func getLiveInput(session string, year, day int) (io.ReadCloser, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", fmt.Sprintf("session=%s", session))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing HTTP request: %w", err)
	}
	return resp.Body, nil
}
