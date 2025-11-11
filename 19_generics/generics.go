package main

import "fmt"

//////////////////////////////////////////////////////////////////////
// 🧩 STEP 1: Define a Generic Function (no restriction)
//////////////////////////////////////////////////////////////////////
//
// `printing[T any]` is a generic function that accepts a slice of any type.
// It prints all elements of that slice.
//
func printing[T any](items []T) {
	for _, item := range items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

//////////////////////////////////////////////////////////////////////
// 🧩 STEP 2: Define a Generic Function with Type Constraints
//////////////////////////////////////////////////////////////////////
//
// Here we restrict `T` to int, string, or bool only.
// This demonstrates *type-safe specialization*.
//
func printing2[T int | string | bool](items []T) {
	for _, item := range items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

//////////////////////////////////////////////////////////////////////
// 🧩 STEP 3: Define a Generic Struct that Holds an Array/Slice
//////////////////////////////////////////////////////////////////////
//
// ✅ Here’s the key change:
// The struct now holds a **slice of T**, not a single value.
// This means it can store multiple elements of any type.
//
// Why a slice?
// Because Go slices are flexible views into arrays —
// so we can easily initialize the struct with an array using arr[:]
//
type Container[T any] struct {
	data []T // 👈 slice (dynamic array) of generic type T
}

// ✅ Method to display all elements stored inside Container
func (c Container[T]) Display() {
	fmt.Println("Stored data inside Container:")
	for _, v := range c.data {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

// ✅ Method to get all data
func (c Container[T]) Get() []T {
	return c.data
}

//////////////////////////////////////////////////////////////////////
// 🧩 STEP 4: MAIN FUNCTION — Demonstration
//////////////////////////////////////////////////////////////////////
func main() {

	//////////////////////////////////////////////////////////////
	// 🔹 1️⃣ Define arrays of different types
	//////////////////////////////////////////////////////////////
	//
	// Arrays are fixed-length sequences.
	// We’ll later convert them to slices using arr[:].
	//
	intArray := [4]int{10, 20, 30, 40}
	stringArray := [3]string{"Go", "Generics", "Power"}
	boolArray := [3]bool{true, false, true}

	//////////////////////////////////////////////////////////////
	// 🔹 2️⃣ Convert arrays to slices
	//////////////////////////////////////////////////////////////
	intSlice := intArray[:]       // full slice of intArray
	stringSlice := stringArray[:] // full slice of stringArray
	boolSlice := boolArray[:]     // full slice of boolArray

	//////////////////////////////////////////////////////////////
	// 🔹 3️⃣ Call generic functions
	//////////////////////////////////////////////////////////////
	fmt.Println("Output of printing (T = any):")
	printing(intSlice)
	printing(stringSlice)
	printing(boolSlice)

	fmt.Println("\nOutput of printing2 (T = int|string|bool):")
	printing2(intSlice)
	printing2(stringSlice)
	printing2(boolSlice)

	//////////////////////////////////////////////////////////////
	// 🔹 4️⃣ Using Generic Structs that hold arrays/slices
	//////////////////////////////////////////////////////////////
	fmt.Println("\nOutput of Generic Struct examples (holding slices):")

	// ✅ Create containers with slices (converted from arrays)
	intContainer := Container[int]{data: intArray[:]}
	stringContainer := Container[string]{data: stringArray[:]}
	boolContainer := Container[bool]{data: boolArray[:]}

	// ✅ Display contents
	intContainer.Display()
	stringContainer.Display()
	boolContainer.Display()

	//////////////////////////////////////////////////////////////
	// 🔹 5️⃣ Accessing data returned from struct method
	//////////////////////////////////////////////////////////////
	fmt.Println("\nAccessing data from intContainer using Get():")
	for _, v := range intContainer.Get() {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

//////////////////////////////////////////////////////////////////////
// 🧩 STEP 5: Summary
//////////////////////////////////////////////////////////////////////
//
// ✅ Generic Functions — `printing` and `printing2`
// ✅ Generic Struct — `Container[T]` holds slice (array) of any type
// ✅ Arrays converted to slices using `arr[:]`
// ✅ Fully type-safe — compile-time checks, no type assertion
//
// 🧩 So in short:
// -----------------------------------------------------------
// - Struct holds `[]T` (array/slice of T)
// - Arrays → slices → passed to generic struct
// - One struct & function works for all data types
//////////////////////////////////////////////////////////////////////
