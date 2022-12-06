package cache

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// GetInput returns a puzzle input from the cache.
func GetInput(year, day int) (io.ReadCloser, error) {
	path, err := puzzleInputPath(year, day)
	if err != nil {
		return nil, fmt.Errorf("error getting input path: %w", err)
	}

	exists, err := fileExists(path)
	if err != nil {
		return nil, fmt.Errorf("error checking if input file exists: %w", err)
	} else if !exists {
		return nil, errCacheMiss
	}

	input, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening input file: %w", err)
	}

	return input, nil
}

// SaveInput saves a puzzle input to the cache.
func SaveInput(year, day int, input io.Reader) error {
	path, err := puzzleInputPath(year, day)
	if err != nil {
		return fmt.Errorf("error getting input path: %w", err)
	}

	dir, _ := filepath.Split(path)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("error creating directories for input: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating input file: %w", err)
	}
	defer f.Close()

	_, err = io.Copy(f, input)
	if err != nil {
		return fmt.Errorf("error writing input file: %w", err)
	}

	return nil
}

func puzzleInputPath(year, day int) (string, error) {
	dayDir, err := puzzleDayDir(year, day)
	if err != nil {
		return "", err
	}

	return filepath.Join(dayDir, "input"), nil
}
