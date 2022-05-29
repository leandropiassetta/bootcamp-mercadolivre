package main

import (
	"fmt"
)

func monthlySalaryPerHourWorked(workedHour, hourValue float64) (float64, error) {
	if workedHour < 80 {
		return 0.0, fmt.Errorf("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")
	}
	payment := workedHour * hourValue

	if payment >= 20000 {
		return payment - (payment * 0.10), nil
	}
	return payment, nil
}

func main() {
	salaryOfMonth, err := monthlySalaryPerHourWorked(80, 50)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(salaryOfMonth)
	}

	salaryOfMonth, err = monthlySalaryPerHourWorked(80, 500)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(salaryOfMonth)
	}
	salaryOfMonth, err = monthlySalaryPerHourWorked(70, 50)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(salaryOfMonth)
	}

}
