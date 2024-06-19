package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for i := 'a'; i <= 'z'; i++ {
		if (i-'a')%2 == 0 {
			z01.PrintRune(i)
		} else {
			z01.PrintRune(i - 'a' + 'A')
		}
	}
	z01.PrintRune('\n')
}
//this programe prints even number letters in lowercase and odd number letters in uppercase
