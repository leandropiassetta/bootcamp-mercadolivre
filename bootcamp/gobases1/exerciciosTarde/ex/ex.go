package exerciciosTarde

import "fmt"

func CountWordSizeAndPrintLetters(word string) {
	wordLenght := len(word)
	fmt.Println("A palavra tem ", wordLenght, "caracteres")

	for _, letter := range word {
		fmt.Println(string(letter))
	}
}
