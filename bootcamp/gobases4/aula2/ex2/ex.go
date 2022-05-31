package main

import (
	"fmt"
	"os"
	"reflect"
)

type client struct {
	file     int
	name     string
	lastName string
	RG       int
	phone    int
	address  string
}

func readFileTXT(fileName string) (string, error) {
	path, _ := os.Getwd()

	data, err := os.ReadFile(path + "ex1/" + fileName + ".txt")

	if err != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado")
	}
	return string(data), nil
}

var clients = []client{}

func registerCustomer(name string, lastName string, RG int, phone int, address string) {

	_, err := readFileTXT("customer")
	if err != nil {
		panic("erro: o arquivo indicado não foi encontrado ou está danificado")
	}

	if id, err := generateID(); err != nil {
		panic(err)
	} else {
		clients = append(clients, client{file: id, name: name, lastName: lastName, RG: RG, phone: phone, address: address})
	}

}

func generateID() (int, error) {

	if reflect.TypeOf(clients).String() != "[]client" {
		panic("interrompa a execução")
	}

	return len(clients) + 1, nil
}

func main() {
	err := recover()

	fmt.Println(err)

	registerCustomer("Leandro", "Piasseta", 5464644, 34758322, "Rua das batata")
	fmt.Println(clients)
}
