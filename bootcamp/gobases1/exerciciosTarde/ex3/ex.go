package main

import (
	"fmt"
)

var (
	number int
	months [12]string
)

func correspondingMonthNumer(number int) (string, error) {
	if number > 12 || number < 1 {
		err := fmt.Errorf("Número inválido")
		// fmt.Println(err.Error())
		return "", err
	}

	months = [12]string{"january", "February ", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	return months[number-1], nil
}

func main() {
	number = 13
	fmt.Println(correspondingMonthNumer(number))
}
