package main

import (
	"github.com/01-edu/z01"
)

func main() {
	word := "Hello Allan!"
	for _, char := range word {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
