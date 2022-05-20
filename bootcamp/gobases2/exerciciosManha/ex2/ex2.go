package main

import "fmt"

func calculateAverage(grades ...int) float64 {
	sumGrade := 0.0

	for _, grade := range grades {
		sumGrade += float64(grade)
	}
	average := sumGrade / float64(len(grades))

	return float64(average)
}

func main() {
	fmt.Println(calculateAverage(10, 5, 10, 5))
}
