package client

import (
	"errors"
	"fmt"
	"os"
)

type cacheResult int

const (
	cacheError cacheResult = iota
	cacheHitCorrect
	cacheHitIncorrect
	cacheMiss
)

func checkAnswerCache(year, day, level int, answer int) (cacheResult, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return cacheError, fmt.Errorf("error finding home directory: %w", err)
	}

	answerFile := fmt.Sprintf("%s/.cache/aoc/year-%d/day-%d/level-%d/answer-%d", homeDir, year, day, level, answer)

	info, err := os.Stat(answerFile)
	if os.IsNotExist(err) {
		return cacheMiss, nil
	}
	if err != nil {
		return cacheError, fmt.Errorf("answer file stat error: %w", err)
	}
	if info.IsDir() {
		return cacheError, errors.New("answer file must not be a directory")
	}

	data, err := os.ReadFile(answerFile)
	if err != nil {
		return cacheError, fmt.Errorf("error reading answer file: %w", err)
	}

	if string(data) == "correct" {
		return cacheHitCorrect, nil
	} else if string(data) == "incorrect" {
		return cacheHitIncorrect, nil
	} else {
		return cacheError, errors.New("invalid answer file data")
	}
}

func updateAnswerCache(year, day, level int, answer int, correct bool) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error finding home directory: %w", err)
	}

	answerDir := fmt.Sprintf("%s/.cache/aoc/year-%d/day-%d/level-%d", homeDir, year, day, level)

	err = os.MkdirAll(answerDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating puzzle answer directory: %w", err)
	}

	answerFile := fmt.Sprintf("%s/answer-%d", answerDir, answer)

	var data []byte
	if correct {
		data = []byte("correct")
	} else {
		data = []byte("incorrect")
	}

	err = os.WriteFile(answerFile, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing answer file: %w", err)
	}

	return nil
}
