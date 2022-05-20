package main

import "fmt"

const (
	categoryA = "a"
	categoryB = "b"
	categoryC = "c"
)

var categorySalaryHour = map[string]int{categoryA: 3000, categoryB: 1500, categoryC: 1000}

type Employee struct {
	name        string
	category    string
	workedHours int
}

func calculateSalary(e Employee) float64 {
	salary := float64(categorySalaryHour[e.category] * e.workedHours)

	if e.category == categoryA && e.workedHours > 160 {
		return (salary * 0.5) + salary
	}

	if e.category == categoryB && e.workedHours > 160 {
		return (salary * 0.2) + salary
	}

	return salary
}

func main() {
	employee_1 := Employee{
		name:        "Leandro de Freitas Piasseta",
		category:    "c",
		workedHours: 160,
	}

	fmt.Println(calculateSalary(employee_1))
}
