package advent

import (
	"errors"
	"fmt"
	"github.com/scjudd/aoc-2022/pkg/advent/internal/cache"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var errAlreadyComplete error = errors.New("already complete")
var errInvalidAnswerType error = errors.New("invalid answer type")

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

func serializeAnswer(v interface{}) (string, error) {
	var s string
	switch v.(type) {
	case int:
		s = strconv.Itoa(v.(int))
	case string:
		s = v.(string)
	default:
		return "", errInvalidAnswerType
	}
	return s, nil
}

func checkAnswer(s *State, level int, answer interface{}) (correct bool, err error) {
	answerString, err := serializeAnswer(answer)
	if err != nil {
		return false, err
	}

	correct, err = cache.GetAnswer(s.Year, s.Day, level, answerString)
	if err == nil {
		return correct, nil
	} else if !cache.IsCacheMiss(err) {
		return false, fmt.Errorf("error checking puzzle answer cache: %w", err)
	}

	correct, err = submitAnswer(s.Session, s.Year, s.Day, level, answerString)
	if err == errAlreadyComplete {
		answerOne, answerTwo, err := getPreviousAnswers(s.Session, s.Year, s.Day)
		if err != nil {
			return false, fmt.Errorf("error fetching previous answers: %w\n", err)
		}
		if answerOne != "" {
			err = cache.SaveAnswer(s.Year, s.Day, 1, answerOne, true)
			if err != nil {
				return false, fmt.Errorf("error updating puzzle answer cache: %w", err)
			}
		}
		if answerTwo != "" {
			err = cache.SaveAnswer(s.Year, s.Day, 2, answerTwo, true)
			if err != nil {
				return false, fmt.Errorf("error updating puzzle answer cache: %w", err)
			}
		}
		if level == 1 && answerString == answerOne {
			return true, nil
		}
		if level == 2 && answerString == answerTwo {
			return true, nil
		}
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("error submitting answer: %w", err)
	}

	err = cache.SaveAnswer(s.Year, s.Day, level, answerString, correct)
	if err != nil {
		return correct, fmt.Errorf("error updating puzzle answer cache: %w", err)
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
