package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "James", lastName: "Jones"}
	// reuse and embed structs
	jim := person{
		firstName: "Jim",
		lastName:  "Hambert",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipcode: 94115,
		},
	}
	jimPointer := &jim
	// & is an operator: give me access memory address of the value this variable is pointing at
	jimPointer.updateName("Jimmy")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) { // * we are looking for a pointer to the type
	//* operator: give me the value this memory address is pointing at
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

/*
Go is a pass by value language
Pointers:
				RAM
Address     Values
0000	->
0001	->  person(firstName: "jim"...)  <- jim
0002	->
0003	->	person(firstName: "Jim"...)  <- p
0004	->
*/
