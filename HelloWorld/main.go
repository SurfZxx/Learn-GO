package main

func newCard() string {
	return "Five of Diamonds"
}

func main() {
	// var i = 2
	// fmt.Println("Hello there, this is user:", i)
	// fmt.Println(time.Now()) //time and date at run moment

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}
