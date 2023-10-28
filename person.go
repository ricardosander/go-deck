package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func newPerson(firstName string, lastName string, email string, zipCode int) person {
	return person{
		firstName:   firstName,
		lastName:    lastName,
		contactInfo: newContactInfo(email, zipCode),
	}
}

func otherNewPerson(firstName string, lastName string, email string, zipCode int) person {
	return person{firstName, lastName, newContactInfo(email, zipCode)}
}

func anotherNewPerson(firstName string, lastName string, email string, zipCode int) person {
	var newPerson person
	newPerson.firstName = firstName
	newPerson.lastName = lastName
	newPerson.contactInfo = newContactInfo(email, zipCode)
	return newPerson
}

func (p *person) updateName(name string) {
	(*p).firstName = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
