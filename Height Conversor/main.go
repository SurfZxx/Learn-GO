package main

import (
	"fmt"
	"strconv"
)

func metrics(feet float64, inch float64, altura float64) float64 {
	altura += feet * 30.48
	altura += inch * 2.54
	return altura
}

func main() {
	feet := "6"
	inch := "4"
	feets, err := strconv.Atoi(feet)
	if err != nil {
		fmt.Println("Error during conversion 1")
		return
	}
	inches, err := strconv.Atoi(inch)
	if err != nil {
		fmt.Println("Error during conversion 2")
		return
	}
	altura := 0
	result := metrics(float64(feets), float64(inches), float64(altura))
	fmt.Println(result)

}
