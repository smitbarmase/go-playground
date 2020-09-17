package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	smit := person{
		firstName: "Smit",
		lastName:  "Barmase",
		contact: contactInfo{
			email:   "smitbarmase@outlook.com",
			zipCode: 456789,
		},
	}
	smit.updateName("Smityo")
	smit.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
