package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func newPerson(firstName string, lastName string, email string, zipCode int) person {
	return person{
		firstName: firstName,
		lastName:  lastName,
		contact:   newContactInfo(email, zipCode),
	}
}

func otherNewPerson(firstName string, lastName string, email string, zipCode int) person {
	return person{firstName, lastName, newContactInfo(email, zipCode)}
}

func anotherNewPerson(firstName string, lastName string, email string, zipCode int) person {
	var newPerson person
	newPerson.firstName = firstName
	newPerson.lastName = lastName
	newPerson.contact = newContactInfo(email, zipCode)
	return newPerson
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
