package main

import "fmt"

func main() {

	// Example 1: Simple if-else
	age := 20
	if age >= 18 {
		fmt.Println("✅ You are an adult.")
	} else {
		fmt.Println("❌ You are not an adult.")
	}

	// Example 2: if-else if-else chain
	score := 75
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else if score >= 50 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// Example 3: Using logical AND (&&)
	marks := 85
	attendance := 90
	if marks >= 80 && attendance >= 75 {
		fmt.Println("🎉 You passed with good marks and attendance.")
	} else {
		fmt.Println("⚠️ You need to improve either marks or attendance.")
	}

	// Example 4: Using logical OR (||)
	temperature := 35
	if temperature > 40 || temperature < 10 {
		fmt.Println("🥵 Extreme weather condition!")
	} else {
		fmt.Println("🌤 Weather is normal.")
	}

	// Example 5: Using NOT (!) operator
	isRaining := false
	if !isRaining {
		fmt.Println("☀️ It's not raining. You can go outside.")
	} else {
		fmt.Println("🌧 It's raining. Stay indoors!")
	}

	// Example 6: if statement with initialization + print number 
	// we can declare a variable inside the if contstruct
	if num := 10; num%2 == 0 {
		fmt.Printf("🔢 Number %d is even.\n", num)
	} else {
		fmt.Printf("🔢 Number %d is odd.\n", num)
	}

	// Example 7: Nested if (if inside if)
	a, b := 5, 10
	if a < b {
		if a%2 == 1 {
			fmt.Println("✅ a is smaller than b and also odd.")
		} else {
			fmt.Println("✅ a is smaller than b but even.")
		}
	} else {
		fmt.Println("❌ a is not smaller than b.")
	}
}
