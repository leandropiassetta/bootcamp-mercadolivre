package main

import (
	"fmt"
)

var salary int

type customError struct {
	message string
}

func (c *customError) Error() string {
	return fmt.Sprintf("erro: %s", c.message)
}

func salaryTax(salaray float64) (string, error) {
	if salaray < 15000 {
		return "", &customError{
			message: "O salário digitado não está dentro do valor mínimo",
		}
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
