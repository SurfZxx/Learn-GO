package main

import "fmt"

func newCard() string {
	return "Ace of Spades"
}

func main() {
	cards := []string{newCard(), newCard()}
	fmt.Println(cards)

	//testArray := []int{}
	var testArray []int
	for i := 0; i < 5; i++ {
		testArray = append(testArray, i)
	}
	//arrlength := len(testArray)
	for i := 0; i < len(testArray); i++ {
		testArray[i] = 5
	}
	fmt.Println(testArray)
}
