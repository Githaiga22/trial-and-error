/*
write a program with 3 variables , one named b of type byte,
one named smallI of type int32 and one named bigI of type uint64.
Asiign each variable the maximum legal value for its type then
add 1 to each variable ..Print out their variables
*/

package main

import "fmt"

func main() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	
	fmt.Println(b + 1)
	fmt.Println(smallI + 1)
	fmt.Println(bigI + 1)
}
