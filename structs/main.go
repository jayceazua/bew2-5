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
	// jim := person{
	// 	firstName: "Jim",
	// 	lastName:  "Hambert",
	// 	contactInfo: contactInfo{
	// 		email:   "jim@gmail.com",
	// 		zipcode: 94115,
	// 	},
	// }
	// jimPointer := &jim
	// & is an operator: give me access memory address of the value this variable is pointing at
	// jimPointer.updateFirstName("Jimmy")
	// jimPointer.updateLastName("Azua")
	// jim.updateFirstName("Jayce")
	// jim.print()
	mySlice := []string{"hi", "there", "how", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)

}

func updateSlice(s []string) {
	s[0] = "bye"
}

/*
func (pointerToPerson *person) updateFirstName(newFirstName string) { // * type description: we are looking for a pointer to the type
	//* operator: give me the value this memory address is pointing at
	(*pointerToPerson).firstName = newFirstName // * an operator: manipulate the value the pointer is referencing
}

func (pointerToPerson *person) updateLastName(newLastName string) {
	(*pointerToPerson).lastName = newLastName
}



func (p person) print() {
	fmt.Printf("%+v", p)
}
*/
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


Turn address into value with *address
Turn value into address with &value

Whenever you pass a int, string, float, or struct into a function,
	Go creates a copy of each argument, and these copies are used inside of the function


	specifies the type of the receiver that the function expects

	When creating a slice, Go will automatically create an array and a structure
		that records the length of the slice, the capacity of the slice, and a reference to the underlying array.

When the slice is created, we get both an array and a data structure describing that array. 
	The data structure describing the array *is* copied before being passed off to the function, but the underlying array is not
*/
