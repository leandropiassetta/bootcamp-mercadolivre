package main

import "fmt"

var (
	number int
	months [12]string
)

func correspondingMonthNumer(number int) string {
	months = [12]string{"january", "February ", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	return months[number-1]
}

func main() {
	number = 5
	fmt.Println(correspondingMonthNumer(number))
}
