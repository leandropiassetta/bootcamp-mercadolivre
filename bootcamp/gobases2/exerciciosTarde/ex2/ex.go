package main

import "fmt"

const (
	small  = "small"
	medium = "medium"
	big    = "big"
)

type Product interface {
	CalculateCost() float64
}

type Ecommerce interface {
	Add(p product)
	Total() float64
}

type store struct {
	products []product
}

type product struct {
	typeProduct string
	name        string
	price       float64
}

func (s *store) Add(p product) {
	s.products = append(s.products, p)
	fmt.Println(s.products)
}

func (s store) Total() float64 {
	total := 0.0

	for _, product := range s.products {
		total += product.price + product.CalculateCost()

	}

	return total
}

func (p product) CalculateCost() float64 {
	if p.typeProduct == small {
		return p.price
	} else if p.typeProduct == medium {
		return p.price * 0.03
	} else {
		return p.price*0.06 + 2500
	}
}

func newProduct(typeProduct string, name string, price float64) product {
	return product{
		typeProduct,
		name,
		price,
	}
}

func newStore() Ecommerce {
	return &store{}
}

func main() {

	magazineLuiza := newStore()

	tv := product{
		typeProduct: big,
		name:        "Samsung",
		price:       9000,
	}

	radio := product{
		typeProduct: small,
		name:        "teste",
		price:       100,
	}

	magazineLuiza.Add(tv)
	magazineLuiza.Add(radio)

	fmt.Println(magazineLuiza.Total())

}
