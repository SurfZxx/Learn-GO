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
	myHand = generateHands(shuffledDeck, counter)
	counter -= 2
	cpuHand = generateHands(shuffledDeck, counter)
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
	gameCards = round1(shuffledDeck, counter)
	counter -= 3
	fmt.Println("Game Cards:")
	for i := 0; i < len(gameCards); i++ {
		fmt.Println(gameCards[i])
	}

	//Do you want to go to next round?
	fmt.Println("Do you want to keep playing:")
	fmt.Scanln(&answer)
	if answer == 0 {
		losePoint()
	}

	//Print the last 2 game cards
	gameCards = round2(shuffledDeck, gameCards, counter)
	fmt.Println("\n")
	fmt.Println("Game Cards:")
	for i := 0; i < len(gameCards); i++ {
		fmt.Println(gameCards[i])
	}

	//Do you want to go to next round?
	fmt.Println("Do you want to keep playing:")
	fmt.Scanln(&answer)
	if answer == 0 {
		for i := 0; i < 2; i++ {
			losePoint()
		}
	}

	// Compare Hands to see the winner
	winner := game(myHand, cpuHand)
	fmt.Println("Winner is: " + winner)
	if winner == "Player" {
		addPoint()
	} else if winner == "CPU" {
		for i := 0; i < 3; i++ {
			losePoint()
		}
	}
}
