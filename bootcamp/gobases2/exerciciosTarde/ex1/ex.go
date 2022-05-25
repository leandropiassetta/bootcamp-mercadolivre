package main

import (
	"fmt"
	"time"
)

// Interface como um conjunto de comportamentos que são esperados de um tipo

type Student struct {
	name          string
	lastName      string
	RG            int
	admissionDate time.Time
}

var students []Student

func (s Student) printDetail() {
	fmt.Printf("Name: %s\nLastname: %s\nRG: %d\nAdmission: %s\n \n", s.name, s.lastName, s.RG, s.admissionDate)
}

func registerStudent(s ...Student) {

	for _, student := range s {
		students = append(students, student)
	}
}

func detailAllStudents() {
	for _, student := range students {
		student.printDetail()
	}
}

func main() {

	s1 := Student{name: "Leandro", lastName: "Piasseta", RG: 4937274, admissionDate: time.Now().UTC()}
	s2 := Student{name: "Joaquim", lastName: "Ferreira", RG: 593474, admissionDate: time.Date(1980, 01, 15, 0, 0, 0, 0, time.UTC)}
	s3 := Student{name: "Clara", lastName: "Baum", RG: 693474, admissionDate: time.Date(1999, 11, 01, 0, 0, 0, 0, time.UTC)}

	registerStudent(s1, s2, s3)
	detailAllStudents()

}
