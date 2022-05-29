package main

import (
	"errors"
	"fmt"
)

var salary int

func salaryTax(salaray float64) (string, error) {
	if salaray < 15000 {
		return "", errors.New("erro: O salário digitado não está dentro do valor mínimo")
	}
	return "Necessário pagamento de imposto.", nil
}

func main() {
	message, err := salaryTax(14000)

	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		panic(err)
	}

	fmt.Println(message)

}
