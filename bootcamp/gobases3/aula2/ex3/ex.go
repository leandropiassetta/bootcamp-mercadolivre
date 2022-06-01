package main

import (
	"fmt"
	"sync"
)

type product struct {
	name     string
	price    float64
	quantity int
}

type service struct {
	name          string
	price         float64
	minutesWorked int
}

type maintenance struct {
	name  string
	price float64
}

var products = []product{
	{
		"tv",
		35,
		10,
	},
	{
		"radio",
		10,
		10,
	},
}

var services = []service{
	{
		"jardinagem",
		50,
		8,
	},
	{
		"Portaria",
		30,
		8,
	},
}

var maintenances = []maintenance{
	{
		"Janelas",
		100,
	},
}

func sumProduct(p []product) float64 {
	total := 0.0

	for _, prod := range p {
		total += (prod.price * float64(prod.quantity))
	}
	fmt.Println(total)
	return total
}

func sumServices(s []service) float64 {
	total := 0.0

	for _, prod := range s {
		if prod.minutesWorked < 30 {
			total += prod.price * 0.5
		} else {
			total += prod.price * float64(prod.minutesWorked/60)
		}
	}
	fmt.Println(total)
	return total
}

func sumMaintenance(m []maintenance) float64 {
	total := 0.0

	for _, maint := range m {
		total += maint.price
	}

	fmt.Println(total)
	return total
}

func main() {
	priceTotal := 0.0
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		priceTotal += sumProduct(products)
		wg.Done()
	}()
	go func() {
		priceTotal += sumServices(services)
		wg.Done()
	}()
	go func() {
		priceTotal += sumMaintenance(maintenances)
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(priceTotal)
}
