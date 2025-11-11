package main

import "fmt"

//////////////////////////////////////////////////////////////////////
// 🧩 Step 1: The Empty Interface — `interface{}`
//////////////////////////////////////////////////////////////////////
//
// 👉 The empty interface `interface{}` means "no methods".
// 👉 Since it requires no methods, **every type** in Go automatically
//    implements it.
//
// 🧠 In simple words: interface{} can hold a value of ANY type.
//
// It’s like a box 🗳️ that can contain anything — int, string, bool, struct, etc.
//
//////////////////////////////////////////////////////////////////////

func main() {

	// 🧩 Step 2: Store different types inside `interface{}`
	var mysteryBox interface{}

	// ✅ You can assign any type to it
	mysteryBox = "Hello Go!"    // string
	describeValue(mysteryBox)   // Type: string, Value: Hello Go!

	mysteryBox = 42             // int
	describeValue(mysteryBox)   // Type: int, Value: 42

	mysteryBox = 3.14           // float64
	describeValue(mysteryBox)   // Type: float64, Value: 3.14

	mysteryBox = true           // bool
	describeValue(mysteryBox)   // Type: bool, Value: true

	//////////////////////////////////////////////////////////////////////
	// 🧩 Step 3: Type Assertion — Extracting the Real Value
	//////////////////////////////////////////////////////////////////////
	//
	// Syntax:
	//    value := i.(Type)
	//
	// It means: “Assume the value inside the interface `i` is of type `Type`,
	// and extract it.”
	//
	// ⚠️ If your assumption is wrong, it causes a **panic** (runtime crash).
	//////////////////////////////////////////////////////////////////////

	mysteryBox = "String value"
	fmt.Println("\n--- Type Assertion Example ---")

	// ❌ UNSAFE: If type doesn’t match, this will panic
	// Uncomment the below line and run to see panic
	// value := mysteryBox.(int)
	// fmt.Println("Value:", value)

	//////////////////////////////////////////////////////////////////////
	// ✅ SAFE VERSION of Type Assertion
	//////////////////////////////////////////////////////////////////////
	//
	// You can prevent panics using the “comma ok” idiom:
	//     value, ok := i.(Type)
	//
	// - If correct type → ok = true, value contains the real data
	// - If wrong type   → ok = false, value = zero value of that type
	//////////////////////////////////////////////////////////////////////

	retrievedInt, ok := mysteryBox.(int)
	if ok {
		fmt.Println("Retrieved int:", retrievedInt)
	} else {
		fmt.Println("❌ Value is not an integer") // this will run
	}

	retrievedString, ok := mysteryBox.(string)
	if ok {
		fmt.Println("✅ Retrieved string:", retrievedString)
	}

	//////////////////////////////////////////////////////////////////////
	// 🧩 Step 4: Type Switch — Handle Multiple Types Dynamically
	//////////////////////////////////////////////////////////////////////
	//
	// Instead of writing multiple type assertions,
	// you can use a type switch to check several types at once.
	//////////////////////////////////////////////////////////////////////

	fmt.Println("\n--- Type Switch Example ---")

	// Let's store something else inside mysteryBox
	mysteryBox = 99.9 // now a float64

	switch value := mysteryBox.(type) {

	case int:
		fmt.Println("It's an integer:", value)

	case string:
		fmt.Println("It's a string:", value)

	case bool:
		fmt.Println("It's a boolean:", value)

	case float64:
		fmt.Println("It's a float64:", value)

	default:
		fmt.Println("Unknown type!")
	}

	//////////////////////////////////////////////////////////////////////
	// 🧩 Step 5: Why Use interface{}?
	//////////////////////////////////////////////////////////////////////
	//
	// ✅ When you don’t know the type ahead of time (e.g., JSON parsing)
	// ✅ When you want a function that can accept ANY type
	// ✅ When dealing with APIs, dynamic inputs, or reflection
	//
	// But: use it carefully! You lose compile-time type safety.
	//////////////////////////////////////////////////////////////////////
}

//////////////////////////////////////////////////////////////////////
// 🧩 Helper Function: describeValue
//////////////////////////////////////////////////////////////////////
//
// This function takes `t interface{}` → it can accept ANY type.
// Inside, we use %T and %v for introspection:
//
// - %T → prints the actual (dynamic) type stored inside
// - %v → prints the actual (dynamic) value
//
//////////////////////////////////////////////////////////////////////

func describeValue(t interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", t, t)
}
