package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
// which is a slice of strings
type deck []string // <-- extends from slice of string
// receiver function
/*
	any variable of type "deck" now gets access to the "print" method
	receivers sets up method on variables we create
*/

func newDeck() deck {
	// init an empty deck
	cards := deck{}
	// suits array
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clovers"}
	// values array
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	// loop through the suits
	for _, suit := range cardSuits {
		// loop through the values
		for _, value := range cardValues {
			// add a new card of value and suit to the deck
			cards = append(cards, value+" of "+suit)
		}
	}
	// return the deck
	return cards
}

func (d deck) print() { // the d is like the keyword; self or this
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()) /* type conversion to byte */, 0666 /*it is the permission for read and write*/)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil { // error handling
		// log the error and entirely quit the program
		fmt.Println("Error:", err)
		// went into the os package so if we do have an error we just exit out completely from our program
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// shuffle function
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	// for each index, card in cards
	for i := range d {
		// generate a random number between 0 and len(cards) - 1
		newPosition := r.Intn(len(d) - 1)
		// swap the current card and the card at cards[randomNumber]
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
