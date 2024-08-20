/*
write a program that declares a constant called value that can be
assigned to both an interger and a floating-point variable.
Assign it to an interger called i and a floating-point varibale
callef f .Print out i and f
*/

package main

import "fmt"

func main() {
	const value = 50
	i := value
	f := float64(value)

	fmt.Printf("The value of i =%d\n", i)
	fmt.Printf("The value of f =%.1f\n", f)
}
