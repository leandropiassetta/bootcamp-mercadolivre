package main

import "fmt"

type client struct {
	name          string
	age           int
	isEmployed    bool
	workingMonths int
	salary        int
	loan          bool
	interestFree  bool
}

func loan(c client) {
	if c.age > 22 && c.isEmployed && c.workingMonths > 12 {
		c.loan = true
	}

	if c.loan && c.salary < 100000 {
		c.interestFree = true
	}

	if !c.loan {
		fmt.Println("Não autorizado a pedir empréstimo")
	} else if !c.interestFree {
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
		salary:        100000,
	}

	loan(client1)
}
