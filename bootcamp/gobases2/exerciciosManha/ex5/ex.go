package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

//dogFunc returns in grams
func dogFood(quantity int) int {
	return quantity * 10000
}

//catFunc returns in grams
func catFood(quantity int) int {
	return quantity * 5000
}

//hamsterFunc returns in grams
func hamsterFood(quantity int) int {
	return quantity * 250
}

//tarantulaFunc returns in grams
func tarantulaFood(quantity int) int {
	return quantity * 150
}

//Animal returns a func for a specific animal
func Animal(specie string) (func(quantity int) int, error) {
	if specie == dog {
		return dogFood, nil

	} else if specie == cat {
		return catFood, nil

	} else if specie == hamster {
		return hamsterFood, nil

	} else if specie == tarantula {
		return tarantulaFood, nil

	} else {
		return nil, errors.New("invalid animal type")
	}
}

func main() {

	dogFunc, _ := Animal("dog")
	catFunc, _ := Animal(cat)
	tarantulaFunc, _ := Animal(tarantula)
	hamsterFunc, _ := Animal(hamster)

	fmt.Printf("A quantidade de alimento do(s) cachorro(s) (em gramas): %d gramas\n", dogFunc(5))
	fmt.Printf("A quantidade de alimento do(s) gatos(s) (em gramas): %d gramas\n", catFunc(2))
	fmt.Printf("A quantidade de alimento da(s) tarantulas(s) (em gramas): %d gramas\n", tarantulaFunc(3))
	fmt.Printf("A quantidade de alimento do(s) hamster(s) (em gramas): %d gramas\n", hamsterFunc(4))

	_, err := Animal("Urubu")
	if err != nil {
		fmt.Println("animal inválido")
	}

}
