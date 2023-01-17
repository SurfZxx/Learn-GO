package main

var myPoints int = 0
var cpuPoints int = 0

func generateHands(deck []string, counter int) []string {
	var result []string
	result = append(result, deck[counter])
	result = append(result, deck[counter-1])
	return result
}

func round1(deck []string, counter int) []string {
	var result []string
	result = append(result, deck[counter])
	result = append(result, deck[counter-1])
	result = append(result, deck[counter-2])
	return result
}

func round2(deck []string, result []string, counter int) []string {
	result = append(result, deck[counter])
	result = append(result, deck[counter-1])
	return result
}

func game(myHand []string, cpuHand []string) string {
	var winner string
	if myHand[0] < cpuHand[0] {
		winner = "CPU"
	} else {
		winner = "Player"
	}
	return winner
}
