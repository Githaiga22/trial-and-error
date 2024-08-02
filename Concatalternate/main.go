package main

import (
	"fmt"
)

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}

func ConcatAlternate(slice1, slice2 []int) []int {
	arr := []int{}
	if len(slice1) == len(slice2) {
		for i := 0; i < len(slice1); i++ {
			arr = append(arr, slice1[i])
			arr = append(arr, slice2[i])
		}
	}
	if len(slice1) > len(slice2) {
		for i := 0; i < len(slice1); i++ {
			if i < len(slice2) {
				arr = append(arr, slice1[i])
				arr = append(arr, slice2[i])
			} else {
				arr = append(arr, slice1[i])
			}
		}
	}
	if len(slice2) > len(slice1) {
		for i := 0; i < len(slice2); i++ {
			if i < len(slice1) {
				arr = append(arr, slice2[i])
				arr = append(arr, slice1[i])
			} else {
				arr = append(arr, slice2[i])
			}
		}
	}
	return arr
}
