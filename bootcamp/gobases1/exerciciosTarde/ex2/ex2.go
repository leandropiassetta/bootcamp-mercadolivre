package main

import "fmt"

type Client struct {
	name           string
	age            int
	isEmployed     bool
	workingMonths  int
	salary         int
	allowLoan      bool
	isInterestFree bool
}

func loan(client Client) {
	if client.age > 22 && client.isEmployed && client.workingMonths > 12 {
		client.allowLoan = true
	}

	if client.allowLoan && client.salary > 100000 {
		client.isInterestFree = true
	}

	if !client.allowLoan {
		fmt.Println("Não autorizado a pedir empréstimo")
	} else if client.isInterestFree {
		fmt.Println("Autorizado a pedir empréstimo sem juros")
	} else {
		fmt.Println("Autorizado a pedir empréstimo com juros")
	}
}

func main() {
	client1 := Client{
		name:          "Leandro Piasseta",
		age:           35,
		isEmployed:    true,
		workingMonths: 13,
		salary:        90000,
	}

	loan(client1)
}
