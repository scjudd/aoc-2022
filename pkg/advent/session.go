package advent

import (
	"errors"
	"os"
)

// LoadSession loads an adventofcode.com session cookie string from the
// environment. Currently this is done by taking the value of the AOC_SESSION
// environment variable. If AOC_SESSION is unset, this function returns an
// error.
func LoadSession() (string, error) {
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return "", errors.New("AOC_SESSION environment variable not set")
	}
	return session, nil
}

// MustLoadSession is the same as LoadSession, except it panics on any errors.
func MustLoadSession() string {
	session, err := LoadSession()
	if err != nil {
		panic(err)
	}
	return session
}
