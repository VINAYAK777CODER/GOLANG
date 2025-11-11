package main

import "fmt"

////////////////////////////////////////////////////////////
// 🌟 FUNCTION PASSED TO FUNCTION IN GO 🌟
//
// ✅ This program explains how to pass a function
//    as a parameter to another function in Go.
// ✅ It also shows how op(a, b) actually works
//    and how the behavior changes depending on which
//    function we pass.
//
// We’ll go step by step, with full explanation in comments.
////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////
// 🧩 STEP 1 — Create a custom type for function
//
// type operation func(int, int) int
//
// This means:
// 👉 "operation" is a type that represents any function
//     which takes two integers and returns one integer.
//
// For example:
// - func add(a int, b int) int
// - func multiply(a int, b int) int
//   Both match this type signature.
////////////////////////////////////////////////////////////
type operation func(int, int) int

////////////////////////////////////////////////////////////
// 🧩 STEP 2 — Create normal functions that match this type
//
// These functions can be passed as arguments because
// their type matches "operation" → func(int, int) int
////////////////////////////////////////////////////////////

// Adds two numbers
func add(x, y int) int {
	return x + y
}

// Multiplies two numbers
func multiply(x, y int) int {
	return x * y
}

// Subtracts two numbers
func subtract(x, y int) int {
	return x - y
}

////////////////////////////////////////////////////////////
// 🧩 STEP 3 — Create a function that ACCEPTS another function
//
// func calculate(a, b int, op operation) int
//
// 👉 Here, "a" and "b" are just integers,
//    and "op" is a FUNCTION that must match the "operation" type.
//
// Inside this function, we will CALL "op" and pass "a" and "b" to it.
//
// So if op = add → op(a, b) = add(a, b)
//    if op = multiply → op(a, b) = multiply(a, b)
//    if op = subtract → op(a, b) = subtract(a, b)
//
// 🧩 So in short:
//
// op(a, b) → calls whatever function you passed as op,
// using a and b as its arguments.
//
// That’s why it’s called a “function parameter” —
// the behavior of calculate() changes based on which function you pass in.
////////////////////////////////////////////////////////////
func calculate(a, b int, op operation) int {
	fmt.Println("Performing operation...")
	result := op(a, b) // calls the function stored in 'op'
	return result
}

////////////////////////////////////////////////////////////
// 🧩 STEP 4 — MAIN FUNCTION
//
// This is where we’ll call "calculate" and pass different
// functions (add, multiply, subtract) to see how behavior changes.
////////////////////////////////////////////////////////////
func main() {

	fmt.Println("========= Example 1: Passing Named Functions =========")

	// Here we pass 'add' function as 'op' parameter.
	// So inside calculate, op(a, b) becomes add(a, b)
	fmt.Println("Add:", calculate(10, 5, add)) // → 10 + 5 = 15

	// Now pass 'multiply' function.
	// So op(a, b) becomes multiply(a, b)
	fmt.Println("Multiply:", calculate(10, 5, multiply)) // → 10 * 5 = 50

	// Now pass 'subtract' function.
	// So op(a, b) becomes subtract(a, b)
	fmt.Println("Subtract:", calculate(10, 5, subtract)) // → 10 - 5 = 5

	fmt.Println("\n========= Example 2: Passing Anonymous (Inline) Function =========")

	// You can also pass an inline (anonymous) function
	// directly instead of a predefined one.
	//
	// This function calculates (x² + y²)
	result := calculate(8, 3, func(x, y int) int {
		return x*x + y*y
	})
	// Inside calculate:
	// op = func(x, y int) int { return x*x + y*y }
	// So op(8, 3) = 8*8 + 3*3 = 64 + 9 = 73
	fmt.Println("Custom inline function result:", result)

	fmt.Println("\n========= Example 3: Function Returning Function + Passing Function =========")

	// Define a function that doubles a number
	double := func(n int) int {
		return n * 2
	}

	// Define another function that accepts a function and returns another function
	// applyTwice(fn) will return a new function that applies fn() twice
	applyTwice := func(fn func(int) int) func(int) int {
		return func(x int) int {
			return fn(fn(x)) // applies the passed function twice
		}
	}

	// Pass our "double" function to applyTwice
	doubleTwice := applyTwice(double)

	// Now call the new function returned
	// So internally: doubleTwice(5) → double(double(5)) = double(10) = 20
	fmt.Println("Applying double twice on 5:", doubleTwice(5)) // Output: 20

	fmt.Println("\n========= SUMMARY =========")
	fmt.Println("✅ Functions can be passed as parameters (like 'op').")
	fmt.Println("✅ You can pass named functions (add, multiply) or inline anonymous ones.")
	fmt.Println("✅ Inside calculate(), op(a, b) means:")
	fmt.Println("   → Call whatever function was passed as 'op' using a and b as arguments.")
	fmt.Println("✅ This allows flexible and reusable code that changes behavior based on the function you pass.")
}

/*
🧩 Step 1: Define double
double := func(n int) int {
	return n * 2
}


Here, double is a function variable.

It takes an integer n and returns n * 2.

✅ Example: double(5) → returns 10.

No function is called yet — just defined.

🧩 Step 2: Define applyTwice
applyTwice := func(fn func(int) int) func(int) int {
	return func(x int) int {
		return fn(fn(x))
	}
}


applyTwice itself is a higher-order function.

It takes a function fn as a parameter.

It returns a new function (that also takes an int x and returns an int).

Inside that returned function:

fn(fn(x))


means → apply the given function fn two times on x.

Still, nothing runs yet — only definitions.

🧩 Step 3: Call applyTwice(double)
doubleTwice := applyTwice(double)


Now something happens:

We call applyTwice, passing our earlier double function as fn.

So inside applyTwice, fn = double.

applyTwice returns a new function:

func(x int) int {
    return double(double(x))
}


That returned function is stored in the variable doubleTwice.

✅ So now doubleTwice is itself a function that expects an integer argument x.

🧩 Step 4: Call doubleTwice(5)
fmt.Println("Applying double twice on 5:", doubleTwice(5))


Now execution really begins step-by-step:

doubleTwice(5) means we are calling the inner anonymous function (the one returned by applyTwice).

func(x int) int {
    return double(double(x))
}


with x = 5.

Inside it:

First double(x) → double(5) → returns 10.

Then double(10) → returns 20.

The result 20 is returned back to fmt.Println.

✅ Output:

Applying double twice on 5: 20

*/
