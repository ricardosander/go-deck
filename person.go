package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func newPerson(firstName string, lastName string) person {
	return person{
		firstName: firstName,
		lastName:  lastName,
	}
}

func otherNewPerson(firstName string, lastName string) person {
	return person{firstName, lastName}
}

func anotherNewPerson(firstName string, lastName string) person {
	var newPerson person
	newPerson.firstName = firstName
	newPerson.lastName = lastName
	return newPerson
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
