package main

import (
	"fmt"
)

type person struct {
	age    int
	health int
	name   string
}

func (p *person) introduce() {
	fmt.Printf("Hi, I am %v.\n", p.name)
}

type personWithFeelings struct {
	*person // IMPORTANT: if this had name, personWithFeelings.introduce() will not work.
	feel    string
}

func main() {
	Mike := &person{
		age:    25,
		health: 100,
		name:   "Mike",
	}

	Jeff := &personWithFeelings{
		person: &person{
			age:    25,
			health: 85,
			name:   "Jeff",
		},
		feel: "sad",
	}

	Mike.introduce()

	// Jeff.introduce() and Jeff.person.introduce() are the same
	// this is called COMPOSITION (conceptually close/alternative to inheritance)
	Jeff.introduce()
	Jeff.person.introduce()
	fmt.Println(Jeff.age == Jeff.person.age)

}
