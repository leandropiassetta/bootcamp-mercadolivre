package main

import (
	"errors"
)

func Divide(num, den int) (int, error) {
	if den == 0 {
		return 0, errors.New("O denominador não pode ser 0")
	}

	return num / den, nil
}
