package main

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
	person := person{}
	person.firstName = firstName
	person.lastName = lastName
	return person
}
