package main

import "fmt"

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	max := a[0]

	for _, numb := range a {
		if numb > max {
			max = numb
		}
	}
	return max
}
