package advent

import (
	"errors"
	"os"
)

type State struct {
	Session string
	Year    int
	Day     int
}

func FromEnv(year, day int) (*State, error) {
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return nil, errors.New("AOC_SESSION environment variable not set")
	}
	return &State{Session: session, Year: year, Day: day}, nil
}

func MustFromEnv(year, day int) *State {
	state, err := FromEnv(year, day)
	if err != nil {
		panic(err)
	}
	return state
}
