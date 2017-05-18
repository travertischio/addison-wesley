// Printing a set of integers, separated by commas
package main

import (
	"fmt"
)

func comma(s string) string {
	r := []rune(s)
	var str string
	for i, count := len(r)-1, 0; i >= 0; i-- {
		if count > 0 && count%3 == 0 {
			str = "," + str
		}
		str = string(r[i]) + str
		count++
	}
	return str
}

func main() {
	fmt.Println(comma("823728424"))
	fmt.Println(comma("22"))
	fmt.Println(comma("2323"))
}
