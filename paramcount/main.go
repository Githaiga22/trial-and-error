package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	count := '0'

	if len(os.Args) > 1 {
		for range os.Args[1:] {
			count++
		}
		z01.PrintRune(count)
	} else {
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')
}
