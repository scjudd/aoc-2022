package advent

import (
	"errors"
	"os"
)

func LoadSession() (string, error) {
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return "", errors.New("AOC_SESSION environment variable not set")
	}
	return session, nil
}

func MustLoadSession() string {
	session, err := LoadSession()
	if err != nil {
		panic(err)
	}
	return session
}
