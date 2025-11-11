package main

import (
	"fmt"
)

//////////////////////////////////////////////////////////////////////
// 🧩 Step 1: Define a small interface — Shape
//////////////////////////////////////////////////////////////////////
//
// 👉 The Shape interface declares one behavior: calculating the Area.
// 👉 Any type that has a method Area() float64 automatically satisfies this interface.
//
type Shape interface {
	Area() float64
}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 2: Define another small interface — Measurable
//////////////////////////////////////////////////////////////////////
//
// 👉 The Measurable interface declares another behavior: calculating the Perimeter.
// 👉 This keeps interfaces small and focused on one behavior.
//
type Measurable interface {
	Perimeter() float64
}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 3: Compose interfaces into a new one — Geometry
//////////////////////////////////////////////////////////////////////
//
// 👉 Geometry is a **composed interface**.
// 👉 It embeds both Shape and Measurable interfaces.
// 👉 So, any type that implements both Area() and Perimeter() methods
//    automatically satisfies Geometry.
//
type Geometry interface {
	Shape        // includes Area() method
	Measurable   // includes Perimeter() method
}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 4: Define a struct — Rectangle
//////////////////////////////////////////////////////////////////////
//
// 👉 Rectangle has width and height fields.
// 👉 We'll implement both Area() and Perimeter() methods on it.
//
type Rectangle struct {
	width, height float64
}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 5: Implement Area() for Rectangle
//////////////////////////////////////////////////////////////////////
//
// 👉 This method calculates area as width * height.
// 👉 Because of this method, Rectangle now satisfies the Shape interface.
//
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 6: Implement Perimeter() for Rectangle
//////////////////////////////////////////////////////////////////////
//
// 👉 This method calculates perimeter as 2*(width + height).
// 👉 Because of this, Rectangle also satisfies Measurable.
//
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 7: Use the composed interface in a function
//////////////////////////////////////////////////////////////////////
//
// 👉 describeShape() accepts any type that implements Geometry.
// 👉 Since Geometry includes both Shape + Measurable,
//    the type must have both Area() and Perimeter() methods.
//
func describeShape(g Geometry) {
	fmt.Println("Area:", g.Area())
	fmt.Println("Perimeter:", g.Perimeter())
}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 8: Main function
//////////////////////////////////////////////////////////////////////
//
// 👉 We create a Rectangle instance and pass it to describeShape().
// 👉 Rectangle implements both Area() and Perimeter(),
//    so it satisfies Geometry.
//
func main() {
	rect := Rectangle{width: 20, height: 10}
	describeShape(rect)
}
