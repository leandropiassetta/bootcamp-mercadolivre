package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(fileName string) {
	path, _ := os.Getwd()

	data, err := os.ReadFile(path + "/gobases3/exercicioDia1/ex1/" + fileName)
	if err != nil {
		panic(err)
	}

	content := strings.Split(string(data), "\n")

	total := 0.0

	fmt.Printf("id\t%25v\t%15v\n", "price", "quantity")

	for _, data := range content[1:] {
		noComma := strings.Split(data, ",")

		id, price, quantity := noComma[0], noComma[1], noComma[2]

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			panic(err)
		}
		convertedQuantity, err := strconv.ParseFloat(quantity, 64)
		if err != nil {
			panic(err)
		}

		total += convertedPrice * convertedQuantity

		fmt.Printf("%v\t%25v\t%15v\n", id, price, quantity)
	}

	fmt.Printf("\t%25v\n", total)
}

func main() {
	readFile("product.csv")
}
