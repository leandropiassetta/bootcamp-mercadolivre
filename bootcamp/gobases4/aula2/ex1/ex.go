package main

import (
	"fmt"
	"os"
)

func readFileTXT(fileName string) (string, error) {
	path, _ := os.Getwd()

	data, err := os.ReadFile(path + "ex1/" + fileName + ".txt")

	if err != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado")
	}
	return string(data), nil
}

func main() {
	customerTxt, err := readFileTXT("customer")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(customerTxt)
}
