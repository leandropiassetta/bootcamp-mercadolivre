package main

import "fmt"

type user struct {
	name     string
	lastName string
	age      int
	email    string
	password string
}

func (u *user) changeName(name string, lastName string) {
	u.name = name
	u.lastName = lastName
}

func (u *user) changeAge(age int) {
	u.age = age
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func (u *user) changePassword(password string) {
	u.password = password
}

func main() {
	user1 := user{
		name:     "leandro",
		lastName: "piasseta",
		age:      35,
		email:    "leandropiasseta@bol.com",
		password: "hahahugf",
	}
	fmt.Println(user1)
	user1.changeEmail("leandropiasseta@UOL.com")
	fmt.Println(user1)
}
