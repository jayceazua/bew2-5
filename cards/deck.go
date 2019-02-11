package main

import "fmt"

// Create a new type of deck
// which is a slice of strings
type deck []string // <-- extends from slice of string
// receiver function
/*
	any variable of type "deck" now gets access to the "print" method
	receivers sets up method on variables we create
*/
func (d deck) print() { // the d is like the keyword; self or this
	for i, card := range d {
		fmt.Println(i, card)
	}
}
