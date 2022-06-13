package main

import (
	"errors"
	"testing"
)

func TestDivide_Error(t *testing.T) {
	num := 10
	den := 0

	expected := errors.New("O denominador não pode ser 0")
	expectedNumber := 0

	res, err := Divide(num, den)

	if err.Error() != expected.Error() {
		t.Errorf("expected: %v, received: %v", expected.Error(), err.Error())
	}

	if res != 0 {
		t.Errorf("expected: %v, received: %v", expectedNumber, res)
	}

}

func TestDivide_Sucess(t *testing.T) {
	num := 10
	den := 2

	expected := 5

	res, err := Divide(num, den)

	if res != expected {
		t.Errorf("expected: %v, received: %v", expected, res)
	}

	if err != nil {
		t.Errorf("expected: %v, received: %v", nil, err.Error())
	}

}
