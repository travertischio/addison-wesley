// Array and slice oprations
package main

import (
	"fmt"
)

func main() {
	sli := []string{"hey", "hi", "hi", "hello", "what"}
	fmt.Println(sli)
	removeAdjDoubles(sli)
	fmt.Println(sli)

	arr := [5]int{3, 4, 5, 2, 5}
	fmt.Println(arr)
	reverse(&arr)
	fmt.Println(arr)

	slice := []int{3, 4, 5, 6, 7, 8, 9, 0, 23, 34, 56}
	fmt.Println(slice)
	slice = rotate(slice, 5)
	fmt.Println(slice)
}

// in place function to eliminate adjacent duplicates in a string slice
func removeAdjDoubles(s []string) {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-2]
		}
	}
}

func reverse(arr *[5]int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// rotate a slice a given number to the left
func rotate(s []int, r int) []int {
	out := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if i+r >= len(s) {
			out[i] = s[(i+r)-len(s)]
		} else {
			out[i] = s[i+r]
		}
	}
	return out
}
