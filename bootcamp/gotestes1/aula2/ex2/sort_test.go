package main

import "testing"

func TestSortIntegers(t *testing.T) {
	numbers := []int{2, 8, 1, 10}

	expected := []int{1, 2, 8, 10}

	result := SortIntegers(numbers)

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("expected: %d, received: %d", expected[i], result[i])
		}
	}

}
