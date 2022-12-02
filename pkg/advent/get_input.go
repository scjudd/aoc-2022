package advent

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetInput(s *State) (io.ReadCloser, error) {
	input, err := checkInputCache(s.Year, s.Day)
	if err != nil {
		return nil, fmt.Errorf("error checking puzzle input cache: %w", err)
	}
	if input != nil {
		return input, nil
	}

	input, err = getLiveInput(s.Session, s.Year, s.Day)
	if err != nil {
		return nil, fmt.Errorf("error getting puzzle input: %w", err)
	}

	err = updateInputCache(s.Year, s.Day, input)
	if err != nil {
		return nil, fmt.Errorf("error updating puzzle input cache: %w", err)
	}

	input, err = checkInputCache(s.Year, s.Day)
	if err != nil {
		return nil, fmt.Errorf("error checking puzzle input cache, after warming: %w", err)
	}

	return input, nil
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

func checkInputCache(year, day int) (io.ReadCloser, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error finding home directory: %w", err)
	}

	inputFile := fmt.Sprintf("%s/.cache/aoc/year-%d/day-%d/input", homeDir, year, day)

	info, err := os.Stat(inputFile)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("puzzle input file stat error: %w", err)
	}
	if info.IsDir() {
		return nil, errors.New("puzzle input file must not be a directory")
	}

	input, err := os.Open(inputFile)
	if err != nil {
		return nil, fmt.Errorf("error opening puzzle input file: %w", err)
	}

	return input, nil
}

func updateInputCache(year, day int, input io.ReadCloser) error {
	defer input.Close()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error finding home directory: %w", err)
	}

	inputDir := fmt.Sprintf("%s/.cache/aoc/year-%d/day-%d", homeDir, year, day)

	err = os.MkdirAll(inputDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating puzzle input directory: %w", err)
	}

	inputFile, err := os.Create(fmt.Sprintf("%s/input", inputDir))
	if err != nil {
		return fmt.Errorf("error creating puzzle input file: %w", err)
	}
	defer inputFile.Close()

	_, err = io.Copy(inputFile, input)
	if err != nil {
		return fmt.Errorf("error writing puzzle input file: %w", err)
	}

	return nil
}
