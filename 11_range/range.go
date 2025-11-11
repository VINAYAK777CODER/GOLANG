package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	for index, value := range nums {
		fmt.Println(value, index)
	}

	// looping in map using range
	m := map[int]int{1: 1, 2: 2, 3: 3}

	// for key only
	for k := range m {
		fmt.Println(k)
	}

	// for key value both
	for k, v := range m {
		fmt.Println(k, v)
	}

	// unicode point rune
	// starting byte of rune
	// 300-> 1byte,2byte
	for i, c := range "golang" {
		fmt.Println(i, c, string(c))

	}
	/*
		i = position in string (starting from 0)

		c = numeric Unicode value of the character

		string(c) converts that rune back to its readable character
	*/

}
