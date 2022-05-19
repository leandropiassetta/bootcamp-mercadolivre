package main

import "fmt"

var (
	nome  string
	idade int
)

func MeusDados() {
	nome = "Leandro de Freitas Piassetta"
	idade = 35

	fmt.Println(nome, ":", idade, "anos")
}
