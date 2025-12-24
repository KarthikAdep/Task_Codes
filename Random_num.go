package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var min, max int
	fmt.Println("Enter the minimum number:")
	fmt.Scan(&min)
	fmt.Println("Enter the maximum number:")
	fmt.Scan(&max)

	var guess int
	var input string

	for {
		
		guess = rand.Intn(max-min+1) + min
		fmt.Println("Is your number", guess, "?")
		fmt.Println("Type 'h' if your number is higher, 'l' if lower, 'c' if correct:")
		fmt.Scan(&input)

		if input == "h" {
			min = guess + 1
		} else if input == "l" {
			max = guess - 1
		} else if input == "c" {
			fmt.Println("Boom! I found the right number")
			break
		} else {
			fmt.Println("Please type 'h', 'l', or 'c'")
		}
	}
}
