package utils

import (
	"testing"
)

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
