package main

import (
	"testing"
)

func TestSub(t *testing.T) {
	a := 5
	b := 2

	expected := 7

	result := Sub(a, b)

	if result != expected {
		t.Errorf("expected: %d, received: %d", expected, result)
	}

}
