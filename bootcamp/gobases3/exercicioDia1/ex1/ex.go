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
	path, _ := os.Getwd()

	header := []byte("id,price,quantity")

	err := os.WriteFile(path+"/gobases3/exercicioDia1/ex1/"+fileName+".csv", header, 0644)

	if err != nil {
		panic(err)
	}
}

func writeDataCSV(fileName string, p Product) {
	infoProducts := fmt.Sprintf("\n%2.d,%.2f,%d", p.id, p.price, p.quantity)
	// os.O_APPEND|os.O_WRONLY|os.O_CREATE -> comando para da append no arquivo...
	path, _ := os.Getwd()

	file, err := os.OpenFile(path+"/gobases3/exercicioDia1/ex1/"+fileName+".csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}

	file.WriteString(infoProducts)
	file.Close()
}

func main() {
	p1 := Product{11, 10.00, 10}
	p2 := Product{11, 10.00, 10}
	p3 := Product{11, 10.00, 10}

	makeFileCSV("product")
	writeDataCSV("product", p1)
	writeDataCSV("product", p2)
	writeDataCSV("product", p3)
	writeDataCSV("product", p3)
	writeDataCSV("product", p3)

}
