package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	arg1 := os.Args[1]
	arg2 := os.Args[2]
	count := 0
	word := 0

	for _, char := range arg1 {
		for i := count; i <= len(arg2)-1; i++ {
			if rune(arg2[i]) == char {
				count = i + 1
				word++
				break
			}
		}
	}
	if word == len(arg1) {
		fmt.Println(arg1)
	}
}
