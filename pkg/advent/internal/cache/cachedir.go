package cache

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func puzzleDayDir(year, day int) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error looking up home directory: %w", err)
	}

	if year < 2015 {
		return "", errors.New("there are no events prior to 2015")
	}

	if day < 1 || day > 25 {
		return "", errors.New("there are 25 days in an event")
	}

	return filepath.Join(homeDir, ".cache", "aoc", fmt.Sprintf("year-%d", year), fmt.Sprintf("day-%d", day)), nil
}

func puzzleLevelDir(year, day, level int) (string, error) {
	dayDir, err := puzzleDayDir(year, day)
	if err != nil {
		return "", err
	}

	return filepath.Join(dayDir, fmt.Sprintf("level-%d", level)), nil
}

func fileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("file stat error: %w", err)
	}
	if info.IsDir() {
		return false, errors.New("found a directory, expected a regular file")
	}
	return true, nil
}
