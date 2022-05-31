package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	maximum = "maximum"
	average = "average"
)

func operation(fn string) (func(...float64) (float64, error), error) {
	switch fn {
	case minimum:
		return getMinValue, nil
	case maximum:
		return getMaxValue, nil
	case average:
		return getAverage, nil
	default:
		return nil, errors.New("Operation not recognized.")
	}
}

// func getMaximum(numbers ...int) int {
// 	sort.Ints(numbers)

// 	return numbers[len(numbers)-1]
// }

func getMaxValue(grades ...float64) (float64, error) {

	if len(grades) == 0 {
		return 0.0, errors.New("input is required")
	}

	maxValue := grades[0]

	for _, grade := range grades {
		if grade > maxValue {
			maxValue = grade
		}
	}

	return maxValue, nil
}

func getMinValue(grades ...float64) (float64, error) {

	if len(grades) == 0 {
		return 0.0, errors.New("input is required")
	}

	minValue := grades[0]

	for _, grade := range grades {
		if grade < minValue {
			minValue = grade
		}
	}

	return minValue, nil
}

func getAverage(grades ...float64) (float64, error) {
	if len(grades) == 0 {
		return 0.0, errors.New("input is required...")
	}

	sum := 0.0

	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades)), nil
}

func main() {

	if minValue, err := operation(minimum); err != nil {
		fmt.Println(err)
	} else {
		if min, err := minValue(2, 3, 3, 4, 10, 2, 4, 5); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(min)
		}
	}

	if maxValue, err := operation(maximum); err != nil {
		fmt.Println(err)
	} else {
		if max, err := maxValue(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(max)
		}
	}

	if avgValue, err := operation(average); err != nil {
		fmt.Println(err)
	} else {
		if avg, err := avgValue(2, 3, 3, 4, 1, 2, 4, 5); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(avg)
		}
	}
}
