package main

import "fmt"

//////////////////////////////////////////////////////////////////////
// 🧩 STEP 1: Define a generic function with two type parameters
//////////////////////////////////////////////////////////////////////
//
// printSlice[T comparable, V string](items []T, name V)
//
// 🧠 Explanation:
// - `T` is a generic type that must be *comparable* (so we can use == or !=).
// - `V` is restricted to *string type* (so we can pass only string values).
// - `items []T` → a slice of any comparable type (like int, string, bool, etc.).
// - `name V` → an additional string parameter.
//
func printSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

//////////////////////////////////////////////////////////////////////
// 🧩 STEP 2: MAIN FUNCTION – Call the generic function
//////////////////////////////////////////////////////////////////////
func main() {
	// Create slices of different types
	nums := []int{10, 20, 30}
	names := []string{"Go", "Lang"}
	bools := []bool{true, false, true}

	// Call the generic function with int slice and a string
	fmt.Println("Printing int slice:")
	printSlice(nums, "Numbers")

	// Call with string slice
	fmt.Println("\nPrinting string slice:")
	printSlice(names, "Words")

	// Call with bool slice
	fmt.Println("\nPrinting bool slice:")
	printSlice(bools, "Booleans")
}
