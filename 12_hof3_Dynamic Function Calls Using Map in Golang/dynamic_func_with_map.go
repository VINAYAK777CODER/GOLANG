package main

import "fmt"

type operation func(int, int) int

func add(a, b int) int      { return a + b }
func subtract(a, b int) int { return a - b }
func multiply(a, b int) int { return a * b }
func divide(a, b int) int {
	if b == 0 {
		fmt.Println("⚠️ Cannot divide by zero!")
		return 0
	}
	return a / b
}

func main() {
	ops := map[string]operation{
		"add":      add,
		"subtract": subtract,
		"multiply": multiply,
		"divide":   divide,
	}

	a, b := 10, 5

	fmt.Println("========= DYNAMIC FUNCTION CALLS USING MAP =========")
	for name, fn := range ops {
		result := fn(a, b)
		fmt.Printf("%s(%d, %d) = %d\n", name, a, b, result)
	}

	fmt.Println("\n========= MANUAL FUNCTION CALLS =========")
	opName := "multiply"
	fmt.Printf("Using key '%s' → %d * %d = %d\n", opName, a, b, ops[opName](a, b))

	opName = "subtract"
	fmt.Printf("Using key '%s' → %d - %d = %d\n", opName, a, b, ops[opName](a, b))

	fmt.Println("\n========= SAFE LOOKUP EXAMPLE =========")
	opName = "modulus"
	if fn, exists := ops[opName]; exists {
		fmt.Println("Result:", fn(a, b))
	} else {
		fmt.Printf("❌ Operation '%s' not found in map.\n", opName)
	}

	ops["power"] = func(x, y int) int {
		result := 1
		for i := 0; i < y; i++ {
			result *= x
		}
		return result
	}

	fmt.Println("\n========= ADDED NEW FUNCTION DYNAMICALLY =========")
	fmt.Printf("power(%d, %d) = %d\n", a, b, ops["power"](a, b))
}
