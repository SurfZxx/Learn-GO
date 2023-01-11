package main

import (
	"fmt"
	"os"
)

func main() {
	fullDeck := getDeck()
	shuffledDeck := shuffle(fullDeck)
	//for i := 0; i < len(shuffledDeck); i++ {
	//	fmt.Println(shuffledDeck[i])
	//}

	// create hands for player and bot
	counter := len(shuffledDeck) - 1
	var myHand []string
	var cpuHand []string
	myHand = append(myHand, shuffledDeck[counter])
	cpuHand = append(cpuHand, shuffledDeck[counter-1])
	counter -= 2

	myHand = append(myHand, shuffledDeck[counter])
	cpuHand = append(cpuHand, shuffledDeck[counter-1])
	counter -= 2
	fmt.Println("My hand:\n" + myHand[0] + ", " + myHand[1] + "\n")
	fmt.Println("CPU Hand:\n" + cpuHand[0] + ", " + cpuHand[1] + "\n")

	//after seeing your hand, do you want to continue?
	answer := 0
	fmt.Println("Do you want to keep playing:")
	fmt.Scanln(&answer)
	if answer == 0 {
		os.Exit(1)
	}

	//print the first 3 cards
	var gameCards []string
	gameCards = append(gameCards, shuffledDeck[counter])
	gameCards = append(gameCards, shuffledDeck[counter-1])
	gameCards = append(gameCards, shuffledDeck[counter-2])
	counter -= 3
	fmt.Println("Game Cards:")
	for i := 0; i < len(gameCards); i++ {
		fmt.Println(gameCards[i])
	}

	//Do you want to go to next round?
	fmt.Println("Do you want to keep playing:")
	fmt.Scanln(&answer)
	if answer == 0 {
		os.Exit(1)
	}

	//Print the last 2 game cards
	gameCards = append(gameCards, shuffledDeck[counter])
	gameCards = append(gameCards, shuffledDeck[counter-1])
	fmt.Println("\n")
	fmt.Println("Game Cards:")
	for i := 0; i < len(gameCards); i++ {
		fmt.Println(gameCards[i])
	}

	//Do you want to go to next round?
	fmt.Println("Do you want to keep playing:")
	fmt.Scanln(&answer)
	if answer == 0 {
		os.Exit(1)
	}

}
