package main

import (
	"fmt"
	"time"
)

func main() {

	// ===============================
	// Example 1: Basic switch
	// ===============================
	day := 3

	switch day {
	case 1:
		fmt.Println("📅 Monday")
	case 2:
		fmt.Println("📅 Tuesday")
	case 3:
		fmt.Println("📅 Wednesday")
	case 4:
		fmt.Println("📅 Thursday")
	case 5:
		fmt.Println("📅 Friday")
	case 6:
		fmt.Println("📅 Saturday")
	case 7:
		fmt.Println("📅 Sunday")
	default:
		fmt.Println("⚠️ Invalid day")
	}

	// ===============================
	// Example 2: Multiple values in a single case
	// ===============================
	day = 6
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("💼 It's a weekday")
	case 6, 7:
		fmt.Println("🎉 It's the weekend!")
	default:
		fmt.Println("❌ Not a valid day number")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println(time.Now().Weekday())
		fmt.Println(" Aaj week day hai aur chutti hai ")
	default:
		fmt.Println(time.Now().Weekday())
		fmt.Println("aaj chutti nahin hai ")
	}

	// ===============================
	// Example 3: Switch without an expression
	// (works like if-else-if)
	// ===============================
	age := 20

	switch {
	case age < 13:
		fmt.Println("👶 You are a child.")
	case age < 20:
		fmt.Println("🧒 You are a teenager.")
	case age < 60:
		fmt.Println("🧑 You are an adult.")
	default:
		fmt.Println("👴 You are a senior citizen.")
	}

	age = 25
	/*
		Go will check which condition becomes true first — exactly like a chain of
		if condition1 { ... } else if condition2 { ... } ...
	*/
	switch true {
	case age < 13:
		fmt.Println("👶 You are a child.")
	case age < 20:
		fmt.Println("🧒 You are a teenager.")
	case age < 60:
		fmt.Println("🧑 You are an adult.")
	default:
		fmt.Println("👴 You are a senior citizen.")
	}

	// ===============================
	// Example 4: Type switch
	// Used to detect variable type at runtime
	// ===============================
	var x interface{} // interface ke karad kisibhi type ki value hun daal sakte hai x ke andar

	x = 3.14 // try changing this to "hello" or true

	/*

		x (type interface{}) — can hold any value, e.g., 3.14.

		x.(type) — used only inside a type-switch to detect the actual concrete type stored in x.

		The switch creates a new variable (e.g., v) that receives the same value but with its concrete type.

		Inside the matching case, v is that value typed correctly (so if x held 3.14, v is 3.14 as float64).

		%v prints the value; %T prints the type.

		Short: x.(type) → finds the type; v → holds the unwrapped value with that type; %v shows the value; %T shows the type.

	*/

	/*

		But if you didn’t write v :=, Go won’t create a new variable holding the typed value.

		You can only check the type, not access the typed value directly.

		Inside the case, you’d still use x (which is of type interface{}), not the unwrapped concrete type.

		🟢 In short:

		switch v := x.(type) → you get both type and typed value in v.

		switch x.(type) → you get only type checking, no typed variable.
	*/

	switch v := x.(type) {
	case int:
		fmt.Printf("🔢 %v is an integer.\n", v)
	case float64:
		fmt.Printf("💧 %v is a float.\n", v)
	case string:
		fmt.Printf("📝 \"%v\" is a string.\n", v)
	case bool:
		fmt.Printf("🔘 %v is a boolean.\n", v)
	default:
		fmt.Println("❓ Unknown type")
	}

	// 🔹 Anonymous function assigned to variable `whoAmI`
	//    - It takes one parameter `i` of type `interface{}`
	//    - interface{} means: i can hold ANY type (int, string, bool, etc.)
	whoAmI := func(i interface{}) { //-----Go me function ke parameter list me var likhna syntax error hai — sirf variable ka naam aur type likhna allowed hai.---//

		// 🔹 Type switch: checks the dynamic type of variable `i`
		switch t := i.(type) { // 👈 't' gets the actual type of 'i'
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("It's a string")
		case bool:
			fmt.Println("It's a boolean")
		default:
			fmt.Println("Other type:", t)
		}
	}

	// 🔹 Function calls with different data types
	whoAmI("golang") // string → prints "It's a string"
	whoAmI(42)       // int → prints "It's an integer"
	whoAmI(true)     // bool → prints "It's a boolean"
	whoAmI(3.14)     // float → goes to default → prints "Other type: 3.14"

	// ===============================
	// Example 5: Switch with fallthrough
	// (manually continues to next case)
	// ===============================
	num := 1
	switch num {
	case 1:
		fmt.Println("Case 1 matched")
		fallthrough // forces the next case to run too
	case 2:
		fmt.Println("Case 2 also executed due to fallthrough")
	default:
		fmt.Println("Default case")
	}
}
