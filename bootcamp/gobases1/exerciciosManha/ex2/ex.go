package main

import "fmt"

var (
	temp     float64
	moisture uint8
	hPa      uint16
)

func main() {
	temp = 14.0
	moisture = 75
	hPa = 1013

	fmt.Println()
	fmt.Printf("Temperatura %v, umidade %v, pressão atmosférica %v hpa", temp, moisture, hPa)
	fmt.Println()
}
