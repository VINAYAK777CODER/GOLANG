package main

import "fmt"

//////////////////////////////////////////////////////
// EXPLANATION: FUNCTION BASICS IN GO
// -----------------------------------
// 1. Go functions can take parameters and return values.
// 2. You can return one or multiple values.
// 3. You can use shorthand type declaration for parameters of same type.
// 4. You can ignore unused return values using `_` (blank identifier).
//////////////////////////////////////////////////////

// ----------- Example 1: Function with single return value -------------
func add(a int, b int) int {
	// Both parameters explicitly have type 'int'
	// Function returns their sum (int)
	return a + b
}

// ----------- Example 2: Shorthand parameter type declaration ----------
func multiply(x, y int) int {
	// When multiple parameters have the same type,
	// you can write the type once after the last parameter.
	return x * y
}

// ----------- Example 3: Function returning multiple values ------------
func operations(a, b int) (int, int, int) {
	// Function returns three values: sum, difference, and product
	sum := a + b
	diff := a - b
	product := a * b
	return sum, diff, product
}

//////////////////////////////////////////////////////
// MAIN FUNCTION
//////////////////////////////////////////////////////
func main() {
	// ---------- Calling function with single return ----------
	result1 := add(5, 3)
	fmt.Println("Sum from add():", result1)

	// ---------- Calling function with shorthand parameters ----------
	result2 := multiply(4, 6)
	fmt.Println("Product from multiply():", result2)

	// ---------- Calling function that returns multiple values ----------
	a, b, c := operations(10, 5)
	fmt.Println("From operations():", a, b, c)

	// ---------- Ignoring unused return value ----------
	// If you only want two values, use '_' (blank identifier)
	sum, _, product := operations(8, 2)
	fmt.Println("From operations() using only 2 values:", sum, product)

	// ---------- Example of ignoring multiple values ----------
	onlySum, _, _ := operations(7, 3)
	fmt.Println("Only sum used:", onlySum)

   

}
