package cache

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const (
	answerCorrect   = "correct"
	answerIncorrect = "incorrect"
)

func GetAnswer(year, day, level int, answer string) (bool, error) {
	path, err := answerPath(year, day, level, answer)
	if err != nil {
		return false, fmt.Errorf("error getting answer path: %w", err)
	}

	exists, err := fileExists(path)
	if err != nil {
		return false, fmt.Errorf("error checking if answer file exists: %w", err)
	} else if !exists {
		return false, errCacheMiss
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return false, fmt.Errorf("error reading answer file: %w", err)
	}

	if string(data) == answerCorrect {
		return true, nil
	} else if string(data) == answerIncorrect {
		return false, nil
	} else {
		return false, errors.New("invalid answer file data")
	}
}

func SaveAnswer(year, day, level int, answer string, correct bool) error {
	path, err := answerPath(year, day, level, answer)
	if err != nil {
		return fmt.Errorf("error getting answer path: %w", err)
	}

	dir, _ := filepath.Split(path)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("error creating directories for answer: %w", err)
	}

	var fileMode os.FileMode = 0644
	if correct {
		err = os.WriteFile(path, []byte(answerCorrect), fileMode)
	} else {
		err = os.WriteFile(path, []byte(answerIncorrect), fileMode)
	}

	if err != nil {
		return fmt.Errorf("error writing answer file: %w", err)
	}

	return nil
}

func answerPath(year, day, level int, answer string) (string, error) {
	levelDir, err := puzzleLevelDir(year, day, level)
	if err != nil {
		return "", err
	}

	return filepath.Join(levelDir, fmt.Sprintf("answer-%s", answer)), nil
}
