package main

func main() { // this function automatically runs
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards /* the slice we want to append */, "Six of Spades") 
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

// go has no idea of classes 
// functions with a receiver 