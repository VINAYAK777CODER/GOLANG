package main

import "fmt"

func main() {

	// -------------------------------------
	// 🌸 Example 1: Capacity grows to 10
	// -------------------------------------
	fmt.Println("===== Example 1: Capacity becomes 10 =====")

	// make() creates a slice with len=2, cap=5
	s1 := make([]int, 2, 5)
	fmt.Println("Initial:", len(s1), "len |", cap(s1), "cap") // 2 | 5

	// Append 3 elements → fits within old capacity (5)
	s1 = append(s1, 1, 2, 3)
	fmt.Println("After 3 appends:", len(s1), "len |", cap(s1), "cap") // 5 | 5

	// Append one more → exceeds capacity → Go reallocates
	s1 = append(s1, 4)
	fmt.Println("After 4th append:", len(s1), "len |", cap(s1), "cap") // 6 | 10

	// ✅ Explanation:
	// oldCap = 5, need = 6
	// Go runtime doubles it → newCap = 10

	// -------------------------------------
	// 🌸 Example 2: Capacity grows to 12
	// -------------------------------------
	fmt.Println("\n===== Example 2: Capacity becomes 12 =====")

	// make() creates slice with len=5, cap=5
	s2 := make([]int, 5)
	fmt.Println("Initial:", len(s2), "len |", cap(s2), "cap") // 5 | 5

	// Append multiple elements one-by-one
	s2 = append(s2, 1)
	s2 = append(s2, 2)
	s2 = append(s2, 3)
	s2 = append(s2, 4)
	fmt.Println("After 4 appends:", len(s2), "len |", cap(s2), "cap")

	// ✅ Explanation:
	// oldCap = 5, total needed = 9
	// Go runtime grows slightly more than double (for optimization)
	// → newCap = 12 (not 10)
	// This helps reduce frequent reallocations for future appends.

	/*

		⚙️ Why both show cap = 12

	Go’s slice growth is not fixed, it’s runtime-dependent — meaning the actual capacity growth can slightly change between Go versions or optimization strategies.

	🧠 What’s happening in your case

	Old capacity = 5

	You append and exceed it

	Go internally calls a function (growslice) to calculate the new capacity

	And for Go 1.21+ (latest) —
	this function uses a more adaptive growth rule, something like this:

	// Simplified internal idea (not exact)
	newCap = oldCap + oldCap/4 + extra_space


	This often results in:

	5 → 12


	instead of strictly doubling (5 → 10).

	*/

}
