package main

import "fmt"

// makeCounter returns a function (a closure)
func makeCounter() func() int {
	count := 0 // 🔒 local variable — only accessible inside makeCounter

	// 👇 This inner function is the closure
	// It "captures" the variable 'count' from its outer function's scope.

	return func() int {
		count++      // ✅ modifies the 'count' variable from outer scope
		return count // ✅ returns the updated count
	}
}

func main() {
	// 👇 create a counter function using closure
	counter := makeCounter()

	// 💡 Each time we call counter(), it remembers the value of count from previous calls
	fmt.Println(counter()) // Output: 1
	fmt.Println(counter()) // Output: 2
	fmt.Println(counter()) // Output: 3

	// 🚀 If we create a new counter, it will have its own "count"
	newCounter := makeCounter()
	fmt.Println(newCounter()) // Output: 1
}
