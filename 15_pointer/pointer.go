package main

import "fmt"


// changeNum takes a pointer to an int ( *int )
// and modifies the value stored at that memory address
func changeNum(n *int) {
	// n holds the address of num (like a reference)
	fmt.Println("Address received:", n)

	// *n means: go to that memory address and change the value there
	*n = 99

	// ✅ num in main() is changed because both share the same memory
}

func main() {
	num := 10 // 🧮 normal integer variable
	fmt.Println("Before change:", num)

	changeNum(&num) // 👈 pass the address of num using &

	fmt.Println("After change:", num)
}
