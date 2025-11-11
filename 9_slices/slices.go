package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used construct in go
// + useful methods
func main() {
	var arr []int            // arr is declared but not initialized, so it’s nil by default.
	fmt.Println(arr == nil)  // true
	arr2 := []int{}          //Here, arr is an empty slice, not nil. its also dynamic way 
	fmt.Println(arr2 == nil) // false

	arr2 = append(arr2, 1, 2, 3, 4, 5)
	fmt.Println(len(arr2), " arr2 ki length")
	fmt.Println(cap(arr2), "  arr2 ki capacity")

	var dynamic_arrays = make([]int, 0, 5)
	fmt.Println(dynamic_arrays == nil) //false
	fmt.Println(len(dynamic_arrays), " is length")
	fmt.Println(cap(dynamic_arrays), " is capacity")

	dynamic_arrays = append(dynamic_arrays, 1)
	dynamic_arrays = append(dynamic_arrays, 2, 5)
	dynamic_arrays = append(dynamic_arrays, 3)
	dynamic_arrays = append(dynamic_arrays, 5)
	dynamic_arrays = append(dynamic_arrays, 4, 9, 10, 11, 12, 13, 14, 15)

	fmt.Println(len(dynamic_arrays), " is new length")
	fmt.Println(cap(dynamic_arrays), " is new capacity")
	fmt.Println(dynamic_arrays)

	var copy_arr = make([]int, len(dynamic_arrays))
	copy(copy_arr, dynamic_arrays) // copy karne ka tarika
	fmt.Println(copy_arr, dynamic_arrays)

	fmt.Println(copy_arr[0:3], " ", copy_arr[:3]," ",copy_arr[3:]) // slices
	fmt.Println(slices.Equal(copy_arr, dynamic_arrays))

	var slices_2d=[][]int{{1,2,3},{4,5,6}}
	fmt.Println(slices_2d)
}
