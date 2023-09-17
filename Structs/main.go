package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email: "test@gmail.com",
			zip:   12345,
		},
	}

	// Pointer
	// alexPointer := &alex

	// Shortcut
	alex.updateName("Jimmy")
	alex.printDetails()
}

func (p person) printDetails() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}
