package client

import (
	"errors"
	"fmt"
	"os"
)

type client struct {
	session string
	year    int
	day     int
}

func New(year, day int) (*client, error) {
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return nil, errors.New("AOC_SESSION environment variable not set")
	}

	return &client{session, year, day}, nil
}

func Must(year, day int) *client {
	c, err := New(year, day)
	if err != nil {
		panic(err)
	}
	return c
}

func (c *client) SubmitPartOne(answer int) (correct bool, err error) {
	return c.submitAnswer(1, answer)
}

func (c *client) SubmitPartTwo(answer int) (correct bool, err error) {
	return c.submitAnswer(2, answer)
}

func (c *client) PrintPartOneResult(answer int) {
	correct, err := c.submitAnswer(1, answer)
	printSubmissionResult(1, answer, correct, err)
}

func (c *client) PrintPartTwoResult(answer int) {
	correct, err := c.submitAnswer(2, answer)
	printSubmissionResult(2, answer, correct, err)
}

func (c *client) submitAnswer(level int, answer int) (correct bool, err error) {
	result, err := checkAnswerCache(c.year, c.day, level, answer)
	if err != nil {
		return false, fmt.Errorf("error checking answer cache: %w", err)
	}

	if result == cacheHitCorrect {
		return true, nil
	} else if result == cacheHitIncorrect {
		return false, nil
	}

	correct, err = submitAnswer(c.session, c.year, c.day, level, answer)
	if err != nil {
		return false, fmt.Errorf("error submitting answer: %w", err)
	}

	err = updateAnswerCache(c.year, c.day, level, answer, correct)
	if err != nil {
		return correct, fmt.Errorf("error updating answer cache: %w", err)
	}

	return correct, nil
}

func printSubmissionResult(level int, answer int, correct bool, err error) {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"

	fmt.Printf("%sPart %d: %d (", colorReset, level, answer)
	if err != nil {
		fmt.Printf("%s%s", colorYellow, err.Error())
	} else if correct {
		fmt.Printf("%sCORRECT! ⭐", colorGreen)
	} else {
		fmt.Printf("%sIncorrect...", colorRed)
	}
	fmt.Printf("%s)\n", colorReset)
}
