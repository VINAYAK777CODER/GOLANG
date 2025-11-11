package main

import (
	"fmt"
	"math"
)

//////////////////////////////////////////////////////////////////////
// 🧩 Step 1: Define an Interface
//////////////////////////////////////////////////////////////////////
//
// 👉 In Go, an *interface* defines **behavior**, not data.
//    It’s like a “contract” that says:
//      “Any type that has these methods will be considered as implementing me.”
//
// 👉 You don’t explicitly say "implements" in Go like in Java or C++.
//    If a type *has the required methods*, it **automatically satisfies** the interface.
//
// For example:
// ----------------------------------------------------------
// type Shape interface {
//	   Area() float64
// }
// ----------------------------------------------------------
// means → “Any type that has a method `Area() float64` is a Shape.”
//////////////////////////////////////////////////////////////////////

type Shape interface {
	Area() float64 // Behavior contract — every Shape must define an Area() method
}

//////////////////////////////////////////////////////////////////////
// 🧱 Step 2: Define Concrete Types (Structs)
//////////////////////////////////////////////////////////////////////
//
// 👉 These are normal structs that hold data.
// 👉 They represent real-world shapes, like Rectangle and Circle.
// 👉 They will each provide their own version of Area().
//
//////////////////////////////////////////////////////////////////////

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 3: Implement the Interface Methods
//////////////////////////////////////////////////////////////////////
//
// 👉 Go automatically checks if a type “implements” an interface.
// 👉 There’s no keyword like "implements" — implementation is *implicit*.
//
// ✅ IMPORTANT RULE:
// If a struct defines **all methods** listed in an interface,
// it automatically becomes an instance of that interface type.
//
// Example:
// If `Shape` requires an `Area()` method,
// and Rectangle has `func (r Rectangle) Area() float64`,
// then Rectangle **is** a Shape.
//
//////////////////////////////////////////////////////////////////////

// Rectangle implements Shape because it has the method Area() float64
func (r Rectangle) Area() float64 {
	// Formula: Area = width × height
	return r.width * r.height
}

// Circle also implements Shape because it defines the same method
func (c Circle) Area() float64 {
	// Formula: Area = π × r²
	return math.Pi * c.radius * c.radius
}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 4: A Function That Accepts the Interface
//////////////////////////////////////////////////////////////////////
//
// 👉 This function accepts a parameter `s` of type `Shape`.
// 👉 That means it can accept *any value* whose type implements Shape.
//
// This is the real power of interfaces: POLYMORPHISM.
//
// The function doesn’t care *what the actual shape is* —
// Circle, Rectangle, Triangle, etc. — it only cares that it can call `Area()`.
//
// Internally, when we pass a Rectangle or Circle:
// ----------------------------------------------------------
//   calculateArea(rect)
// ----------------------------------------------------------
// Go checks:
//   “Does Rectangle have Area() float64?”
// ✅ Yes → automatically converts rect into an interface value of type Shape.
//
// That’s why you can pass structs directly to a function expecting an interface.
//////////////////////////////////////////////////////////////////////

func calculateArea(s Shape) float64 {
	// 👉 Here, `s` is an interface value.
	//    Internally, it stores two things:
	//    1️⃣ The dynamic type (actual type like Rectangle or Circle)
	//    2️⃣ The dynamic value (the data inside that type)
	//
	// When you call s.Area(), Go automatically dispatches to
	// the correct method implementation (Rectangle.Area or Circle.Area).
	//
	// This is called DYNAMIC DISPATCH — a form of runtime polymorphism.
	//
	return s.Area()

}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 5: MAIN Function — Using the Interface in Action
//////////////////////////////////////////////////////////////////////
//
// 👉 Now we’ll create specific shapes and pass them to our function.
// 👉 Notice that `calculateArea()` expects `Shape`, not Rectangle or Circle.
//    Yet we can pass those structs directly — because Go does the conversion
//    automatically if they implement the interface.
//
//////////////////////////////////////////////////////////////////////

func main() {

	// 🟢 Create specific shape instances (normal structs)
	rect := Rectangle{width: 20, height: 10}
	circ := Circle{radius: 5.0}

	//-----------------------------------------------------------------
	// ✅ Why can we pass rect and circ directly to calculateArea() ?
	//-----------------------------------------------------------------
	// Because Go checks if their types implement the Shape interface.
	//
	// Rectangle has Area() → satisfies Shape
	// Circle has Area() → satisfies Shape
	//
	// So Go automatically converts:
	//    rect (Rectangle) → Shape (interface)
	//    circ (Circle) → Shape (interface)
	//
	// Internally, it’s like:
	//    var s Shape = rect
	//    calculateArea(s)
	//-----------------------------------------------------------------

	// Call calculateArea() with Rectangle
	fmt.Println("Rectangle Area:", calculateArea(rect))

	// Call calculateArea() with Circle

	fmt.Println("circle Area:", calculateArea(circ))

}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 6: Internal Working of Interface (Deep Dive)
//////////////////////////////////////////////////////////////////////
//
// When you call `calculateArea(rect)`:
//
// Go internally wraps your struct in an **interface value**:
//
//    var s Shape = rect
//
// Now, `s` holds two pieces of data internally:
//
//    ┌────────────────────────────────────┐
//    │  Interface value (type Shape)      │
//    ├────────────────────────────────────┤
//    │  Dynamic Type: Rectangle           │
//    │  Dynamic Value: {width: 20, height: 10} │
//    └────────────────────────────────────┘
//
// When you call `s.Area()`, Go looks up the method for the actual dynamic type
// (Rectangle) and calls Rectangle.Area().
//
// This is **runtime polymorphism** — one interface, multiple concrete behaviors.
//
//////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////
// 🧩 Step 7: What if a Struct Doesn’t Implement the Interface?
//////////////////////////////////////////////////////////////////////
//
// Example:
//
// type Triangle struct { base, height float64 }
//
// func main() {
//     t := Triangle{base: 10, height: 5}
//     calculateArea(t) // ❌ Compile-time error
// }
//
// Go will say:
//
//     cannot use t (variable of type Triangle) as Shape value in argument to calculateArea:
//     Triangle does not implement Shape (missing method Area)
//
// ✅ This is a *compile-time* error — so Go ensures type safety
// even when using flexible interfaces.
//
//////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////
// 🧩 Step 8: Why Interfaces are Powerful in Go
//////////////////////////////////////////////////////////////////////
//
// - ✅ They allow different types to share behavior (like inheritance)
// - ✅ They achieve polymorphism in a simple, type-safe way
// - ✅ They reduce coupling — you depend on behavior, not concrete types
// - ✅ Implementation is implicit — no need to declare “implements”
//
// Real-life analogy:
// ----------------------------------------------------------
// Interface → a contract (like a “Shape must have Area()” rule)
// Rectangle → follows that rule
// Circle → also follows that rule
// So both can be treated as Shapes!
// ----------------------------------------------------------
//
// This makes your code modular, reusable, and easy to extend.
//
//////////////////////////////////////////////////////////////////////
