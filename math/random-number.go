package main

import (
	"fmt"
	"math/rand"
)

func getRandomNumber(max int) int {
	return rand.Intn(max)
}

func main() {
	max := 2023
	
	fmt.Print("Um número aleatório: ")
	fmt.Print(getRandomNumber(max))
}
