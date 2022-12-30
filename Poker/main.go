package main

import (
	"fmt"
)

func main() {
	fullDeck := getDeck()
	shuffledDeck := shuffle(fullDeck)
	for i := 0; i < len(shuffledDeck); i++ {
		fmt.Println(shuffledDeck[i])
	}
}
