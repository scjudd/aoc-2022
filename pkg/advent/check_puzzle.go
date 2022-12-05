package advent

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var errAlreadyComplete error = errors.New("already complete")

type PuzzleResult struct {
	Level   int
	Answer  interface{}
	Correct bool
}

func CheckPartOne(s *State, answer interface{}) (*PuzzleResult, error) {
	level := 1
	correct, err := checkAnswer(s, level, answer)
	return &PuzzleResult{Level: level, Answer: answer, Correct: correct}, err
}

func CheckPartTwo(s *State, answer interface{}) (*PuzzleResult, error) {
	level := 2
	correct, err := checkAnswer(s, level, answer)
	return &PuzzleResult{Level: level, Answer: answer, Correct: correct}, err
}

func PrintResult(result *PuzzleResult, err error) {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"

	fmt.Printf("%sPart %d: %v (", colorReset, result.Level, result.Answer)
	if err != nil {
		fmt.Printf("%s%s", colorYellow, err.Error())
	} else if result.Correct {
		fmt.Printf("%sCORRECT! ⭐", colorGreen)
	} else {
		fmt.Printf("%sIncorrect...", colorRed)
	}
	fmt.Printf("%s)\n", colorReset)
}

func checkAnswer(s *State, level int, answer interface{}) (correct bool, err error) {
	result, err := checkAnswerCache(s.Year, s.Day, level, answer)
	if err != nil {
		return false, fmt.Errorf("error checking answer cache: %w", err)
	}

	if result == cacheHitCorrect {
		return true, nil
	} else if result == cacheHitIncorrect {
		return false, nil
	}

	correct, err = submitAnswer(s.Session, s.Year, s.Day, level, answer)
	if err == errAlreadyComplete {
		answerOne, answerTwo, err := getPreviousAnswers(s.Session, s.Year, s.Day)
		if err != nil {
			return false, fmt.Errorf("error fetching previous answers: %w\n", err)
		}
		if answerOne != "" {
			err = updateAnswerCache(s.Year, s.Day, 1, answerOne, true)
			if err != nil {
				return false, fmt.Errorf("error updating answer cache: %w", err)
			}
		}
		if answerTwo != "" {
			err = updateAnswerCache(s.Year, s.Day, 2, answerTwo, true)
			if err != nil {
				return false, fmt.Errorf("error updating answer cache: %w", err)
			}
		}
		if level == 1 && fmt.Sprintf("%v", answer) == answerOne {
			return true, nil
		}
		if level == 2 && fmt.Sprintf("%v", answer) == answerTwo {
			return true, nil
		}
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("error submitting answer: %w", err)
	}

	err = updateAnswerCache(s.Year, s.Day, level, answer, correct)
	if err != nil {
		return correct, fmt.Errorf("error updating answer cache: %w", err)
	}

	return correct, nil
}

func submitAnswer(session string, year, day, level int, answer interface{}) (correct bool, err error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day)
	payload := fmt.Sprintf("level=%d&answer=%v", level, answer)

	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	req.Header.Add("Cookie", fmt.Sprintf("session=%s", session))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("error performing HTTP request: %w", err)
	}
	defer resp.Body.Close()

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("error reading HTTP response body: %w", err)
	}

	respBody := string(respBodyBytes)

	if strings.Contains(respBody, "That's the right answer") {
		return true, nil
	}

	if strings.Contains(respBody, "That's not the right answer") {
		return false, nil
	}

	if strings.Contains(respBody, "Did you already complete it") {
		return false, errAlreadyComplete
	}

	if strings.Contains(respBody, "You gave an answer too recently") {
		r := regexp.MustCompile("You have (.+) left to wait")
		match := r.FindStringSubmatch(respBody)
		return false, fmt.Errorf("rate-limited, %s left", match[1])
	}

	return false, fmt.Errorf("unexpected response: HTTP %d: %s", resp.StatusCode, respBody)
}

type cacheResult int

const (
	cacheError cacheResult = iota
	cacheHitCorrect
	cacheHitIncorrect
	cacheMiss
)

func checkAnswerCache(year, day, level int, answer interface{}) (cacheResult, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return cacheError, fmt.Errorf("error finding home directory: %w", err)
	}

	answerFile := fmt.Sprintf("%s/.cache/aoc/year-%d/day-%d/level-%d/answer-%v", homeDir, year, day, level, answer)

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

func updateAnswerCache(year, day, level int, answer interface{}, correct bool) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error finding home directory: %w", err)
	}

	answerDir := fmt.Sprintf("%s/.cache/aoc/year-%d/day-%d/level-%d", homeDir, year, day, level)

	err = os.MkdirAll(answerDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating puzzle answer directory: %w", err)
	}

	answerFile := fmt.Sprintf("%s/answer-%v", answerDir, answer)

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

func getPreviousAnswers(session string, year, day int) (string, string, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", fmt.Sprintf("session=%s", session))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("error performing HTTP request: %w", err)
	}
	defer resp.Body.Close()

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("error reading HTTP response body: %w", err)
	}

	respBody := string(respBodyBytes)

	re := regexp.MustCompile("Your puzzle answer was <code>(.+?)</code>")
	matches := re.FindAllStringSubmatch(respBody, -1)

	var puzzleOne, puzzleTwo string

	if len(matches) >= 1 {
		puzzleOne = matches[0][1]
	}

	if len(matches) == 2 {
		puzzleTwo = matches[1][1]
	}

	return puzzleOne, puzzleTwo, nil
}
