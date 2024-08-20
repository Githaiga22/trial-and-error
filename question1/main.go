/*
write a program that declares an interger variable
called i with the value 50 .Assign i to a floating-point
variable named f .Print out f and i.
*/

package main

import "fmt"

func main() {
	i := 50
	f := float64(i)
	fmt.Printf("the value of i:=%d\n", i)
	fmt.Printf("the value of f :=%.1f\n", f)
}
