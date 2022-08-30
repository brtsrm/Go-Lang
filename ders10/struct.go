package main

import (
	"fmt"
)

type Human struct {
	firstName string
	lastName  string
	age       int
}

func newHuman() *Human {
	h := new(Human)
	return h
}
func newHuman2(firstname, lastname string, age int) *Human {
	h := new(Human)
	h.firstName = firstname
	h.lastName = lastname
	h.age = age
	return h
}

func main() {
	/*
		a := Human{firstName: "Berat", lastName: "Sarmış", age: 25}

		fmt.Println(a.firstName, a.lastName, a.age)

		b := new(Human)
		b.firstName = "Berat"
		b.lastName = "Sarmış"
		b.age = 25
		fmt.Println(b.firstName, b.lastName, b.age)
		newHumanCrate := newHuman()
		newHumanCrate.firstName = "berat"
		newHumanCrate.lastName = "berat"
		newHumanCrate.age = 12
	*/

	a := newHuman2("Berat", "Sarmış", 25)
	fmt.Println(a.firstName, a.lastName, a.age)

}
