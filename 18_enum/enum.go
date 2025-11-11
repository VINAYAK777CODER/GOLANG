package main

import "fmt"

// ✅ Step 1: Define a new custom type called 'Direction'.
// This is based on int, but it’s now treated as a distinct type.
// So, Direction is a “strongly typed” integer used for our enum.
type Direction int

// ✅ Step 2: Define constants using iota.
// iota automatically increments values starting from 0.
const (
	North Direction = iota // North = 0
	East                   // East = 1
	South                  // South = 2
	West                   // West = 3
)

// ✅ Step 3: Define a String() method for the Direction type.
// Any type that has a String() string method automatically satisfies
// the 'fmt.Stringer' interface in Go.
//
// When you print using fmt.Println(), Go checks if the value implements
// String() method — if yes, it calls it automatically.
//
// Example:
// fmt.Println(dir) → internally calls dir.String() → returns a string name.
func (d Direction) String() string {
	switch d {
	case North:
		return "North" // if value == 0 → return "North"
	case East:
		return "East"  // if value == 1 → return "East"
	case South:
		return "South" // if value == 2 → return "South"
	case West:
		return "West"  // if value == 3 → return "West"
	default:
		return "Unknown" // handles invalid or undefined enum values
	}
}

func main() {
	// ✅ Step 4: Create a variable of type Direction and assign a constant.
	dir := South // South = 2

	// ✅ Step 5: Print the direction.
	// Since Direction has a String() method, fmt.Println() will automatically
	// call dir.String() and print "South" instead of just the number 2.
	fmt.Println("Direction:", dir)

	// ✅ Step 6: Test printing an invalid value.
	// Direction(99) means we are manually creating an invalid enum.
	// The String() method handles it by returning "Unknown".
	invalidDir := Direction(99)
	fmt.Println("Invalid Direction:", invalidDir)
}

/*
🧠 Summary of Workflow:

1️⃣  'type Direction int' creates a new custom type for safety and clarity.
2️⃣  'const (...) = iota' automatically generates sequential values.
3️⃣  'String()' adds human-readable meaning to those numeric values.
4️⃣  'fmt.Println()' automatically calls the String() method if it exists.
5️⃣  Output:
     Direction: South
     Invalid Direction: Unknown

✅ Real-world use:
   - You can use this enum type for things like direction, order status,
     HTTP status codes, or any scenario where fixed named constants are needed.
*/
