package main

import "fmt"

var (
	nome  string
	idade int
)

func main() {
	nome = "Leandro de Freitas Piassetta"
	idade = 35

	fmt.Println(nome, ":", idade, "anos")
}
