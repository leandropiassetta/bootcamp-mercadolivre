package main

import (
	"fmt"
	"os"
)

type Product struct {
	id       int
	price    float64
	quantity int
}

func makeFileCSV(fileName string) {
	header := []byte("id,price,quantity")

	err := os.WriteFile("./bootcamp/gobases3/exercicioDia1/ex1/"+fileName+".csv", header, 0644)

	if err != nil {
		panic(err)
	}
}

func writeDataCSV(fileName string, p Product) {
	infoProducts := fmt.Sprintf("\n%2.d,%.2f,%d", p.id, p.price, p.quantity)

	file, err := os.OpenFile(fileName+".csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}

	file.WriteString(infoProducts)
	file.Close()
}

func main() {
	p1 := Product{11, 10.00, 10}
	makeFileCSV("product")
	writeDataCSV("./bootcamp/gobases3/exercicioDia1/ex1/product", p1)
}
