package main

import "fmt"

type client struct {
	name           string
	age            int
	isEmployed     bool
	workingMonths  int
	salary         int
	allowLoan      bool
	isInterestFree bool
}

func loan(c client) {
	if c.age > 22 && c.isEmployed && c.workingMonths > 12 {
		c.allowLoan = true
	}

	if c.allowLoan && c.salary > 100000 {
		c.isInterestFree = true
	}

	if !c.allowLoan {
		fmt.Println("Não autorizado a pedir empréstimo")
	} else if c.isInterestFree {
		fmt.Println("Autorizado a pedir empréstimo sem juros")
	} else {
		fmt.Println("Autorizado a pedir empréstimo com juros")
	}
}

func main() {
	client1 := client{
		name:          "Leandro Piasseta",
		age:           35,
		isEmployed:    true,
		workingMonths: 13,
		salary:        90000,
	}

	loan(client1)
}
