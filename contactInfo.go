package main

type contactInfo struct {
	email string
	zip   int
}

func newContactInfo(email string, zipCode int) contactInfo {
	return contactInfo{
		email: email,
		zip:   zipCode,
	}
}
