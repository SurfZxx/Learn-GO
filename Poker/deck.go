package main

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func getDeck() []string {
	var deck []string
	file, _ := os.Open("deck.txt")
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		deck = append(deck, fileScanner.Text())
	}
	file.Close()
	return deck
}

func shuffle(fullDeck []string) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(fullDeck), func(i, j int) {
		fullDeck[i], fullDeck[j] = fullDeck[j], fullDeck[i]
	})
	return fullDeck
}
