package main

import (
	"github.com/01-edu/z01"
)

func main() {
	word := "Hello world!"
	for _, char := range word {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
