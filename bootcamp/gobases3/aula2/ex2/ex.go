package main

import "fmt"

type user struct {
	name     string
	lastName string
	email    string
	products []product
}

type product struct {
	name     string
	price    float64
	quantity int
}

func (p *product) newProduct(name string, price float64) {
	p.name = name
	p.price = price
}

func (p *product) addProduct(u *user, prod product, quantity int) {
	p.quantity = quantity
	u.products = append(u.products, *p)
}

func (p *product) deleteProduct(u *user) {
	u.products = make([]product, 0)
	//u.products = nil
}

func main() {
	user1 := user{
		name:     "Leandro",
		lastName: "Piasseta",
		email:    "leandrop@gmail.com",
		products: []product{},
	}

	user2 := user{"Joana", "Faria", "jf@gmail.com", nil}

	tv := product{
		"Samsung",
		10000,
		5,
	}

	fmt.Println(user1.products)

	radio := product{}

	radio.newProduct("Radio do zé", 20.5)

	fmt.Println(radio)

	radio.addProduct(&user2, radio, 10)
	fmt.Println(radio)
	tv.addProduct(&user1, tv, 10)
	fmt.Println(tv)

	fmt.Println(user1)

}
