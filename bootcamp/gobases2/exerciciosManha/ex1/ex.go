package main

import "fmt"

func calculateEmployeeTax(salary float64) float64 {
	if salary < 50000 {
		return salary * 0.17
	} else {
		return salary * 0.1
	}
}

func main() {
	fmt.Printf("%.2f\n", calculateEmployeeTax(45000))
	fmt.Printf("%.2f\n", calculateEmployeeTax(150000))
}
