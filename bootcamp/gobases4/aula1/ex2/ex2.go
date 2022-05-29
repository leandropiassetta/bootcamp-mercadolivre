package main

import (
	"fmt"
)

var salary int

func salaryTax(salaray float64) (string, error) {
	if salaray < 15000 {
		return "", fmt.Errorf("erro: O mínimo tributável é 15000.00 e o salário informado é:[%.2f]", salaray)
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
