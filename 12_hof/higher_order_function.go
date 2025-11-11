package main

import "fmt"

////////////////////////////////////////////////////////////
// 🌟 FUNCTION RETURNING FUNCTION IN GO — BOTH VERSIONS 🌟
//
// This program demonstrates both styles of returning functions:
// 1️⃣ Unnamed parameter version: func(int) int
// 2️⃣ Named parameter version: func(x int) int
//
// It also shows how to:
// - Store returned functions in variables
// - Pass parameters to them
// - Call them directly without storing
// - See closure behavior in action
////////////////////////////////////////////////////////////

///////////////////////////
// VERSION 1: Unnamed parameter
// -----------------------------
// Here we do NOT name the parameter in the returned function type.
// func multiplierV1(factor int) func(int) int
///////////////////////////

func multiplierV1(factor int) func(int) int {
	return func(x int) int {
		// Inner function multiplies 'x' by outer variable 'factor'
		return x * factor
	}
}

///////////////////////////
// VERSION 2: Named parameter
// -----------------------------
// Here we NAME the parameter in the returned function type.
// func multiplierV2(factor int) func(x int) int
// This improves readability but behaves exactly the same.
///////////////////////////

func multiplierV2(factor int) func(x int) int {
	return func(x int) int {
		// Inner function does the same work as Version 1
		return x * factor
	}
}

////////////////////////////////////////////////////////////
// MAIN FUNCTION
////////////////////////////////////////////////////////////
func main() {

	fmt.Println("========= VERSION 1: Unnamed Parameter =========")

	// STEP 1: Store returned functions in variables
	double1 := multiplierV1(2)   // returns func(int) int
	triple1 := multiplierV1(3)   // returns func(int) int
	quadruple1 := multiplierV1(4)

	// STEP 2: Use stored function variables by passing values
	fmt.Println("Double of 5:", double1(5))        // 10
	fmt.Println("Triple of 5:", triple1(5))        // 15
	fmt.Println("Quadruple of 5:", quadruple1(5))  // 20

	// STEP 3: Call directly without storing in variable
	fmt.Println("Direct call (5 * 10):", multiplierV1(10)(5)) // 50

	fmt.Println("\n========= VERSION 2: Named Parameter =========")

	// STEP 1: Store returned functions in variables
	double2 := multiplierV2(2)   // returns func(x int) int
	triple2 := multiplierV2(3)   // returns func(x int) int
	quadruple2 := multiplierV2(4)

	// STEP 2: Use stored function variables by passing values
	fmt.Println("Double of 10:", double2(10))       // 20
	fmt.Println("Triple of 10:", triple2(10))       // 30
	fmt.Println("Quadruple of 10:", quadruple2(10)) // 40

	// STEP 3: Call directly without storing in variable
	fmt.Println("Direct call (10 * 5):", multiplierV2(5)(10)) // 50

	fmt.Println("\n========= SUMMARY =========")
	fmt.Println("Version 1 and Version 2 behave identically.")
	fmt.Println("The only difference is that Version 2 names the parameter 'x' in the return type for clarity.")
}
