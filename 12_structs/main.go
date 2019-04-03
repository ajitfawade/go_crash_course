package main

import (
	"fmt"
	"strconv"
)

// Define Person struct

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int

	// alternative
	// firstName, lastName, city, gender string
	// age int
}

// Greeting method (value receiver)

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (Pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (Pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct

	// person1 := Person{firstName: "Ajit", lastName: "Fawade", city: "Pune", gender: "Male", age: 29}
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "F", age: 25}
	// fmt.Println(person1)

	// person1.age++
	// fmt.Println(person1.age)

	person1.hasBirthday()
	person1.getMarried("Williams")
	fmt.Println(person1.greet())

	// Alternative

	// person1 := Person{"Ajit", "Fawade", "Pune", "Male", 29}
	// fmt.Println(person1)
}
