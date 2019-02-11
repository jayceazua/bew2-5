package main

func main() { // this function automatically runs
	/*
	We created two new references that point at subsections of the 'cards' slice.
	We never directly modified the slice that 'cards' is pointing at.
	*/
	cards := newDeck()
	cards.saveToFile("my_cards")
	// type conversion
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
}



// go has no idea of classes
// functions with a receiver
