package utils

import (
	"os"
	"testing"
)

var DEV = false

func init() {
	if len(os.Getenv("DEV")) > 0 {
		DEV = true
	}
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func TPanic(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
