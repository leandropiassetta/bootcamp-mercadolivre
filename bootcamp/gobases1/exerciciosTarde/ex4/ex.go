package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

func main() {
	fmt.Println("A idade de Benjamin é:", employees["Benjamin"])

	var employeesMore21Years []string

	defer func() {
		for key, value := range employees {
			if value > 21 {
				employeesMore21Years = append(employeesMore21Years, key)
			}
		}
		fmt.Println("O número de funcionários acima de 21 anos é:", len(employeesMore21Years))
	}()

	employees["Frederico"] = 25
	delete(employees, "Pedro")
}
