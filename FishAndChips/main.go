package main

import (
	"fmt"
)

func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
	fmt.Println(FishAndChips(55))
	fmt.Println(FishAndChips(-9))
}

func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}
	if n%2 == 0 && n%3 == 0 {
		return "Fish and Chips"
	}
	if n%2 == 0 {
		return "Fish"
	}
	if n%3 == 0 {
		return "Chips"
	}
	
	return "error: non divisible"
}
