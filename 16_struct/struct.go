package main

import (
	"fmt"
	"time"
)

//////////////////////////////////////////////////////////////////////
// 🧱 STEP 1: Define a struct (Custom Data Type)
//////////////////////////////////////////////////////////////////////
//
// 👉 In Go, a `struct` is a collection of fields (like an object template).
// 👉 Here, we're defining a struct named `Order` with 4 fields.
//
type Order struct {
	id      int
	amount  float64
	status  string
	created time.Time
}

//////////////////////////////////////////////////////////////////////
// 🧩 STEP 2: Define Methods with Receivers
//////////////////////////////////////////////////////////////////////

// 🔹 Receiver concept:
// A "receiver" is like a function parameter that allows a function
// to be tied to a specific struct type.
//
// Syntax:
//     func (receiverName ReceiverType) methodName(params...) returnType { ... }
//
// There are two kinds of receivers:
// 1️⃣ Value Receiver  ->  (o Order)     → Works on a COPY of the struct
// 2️⃣ Pointer Receiver ->  (o *Order)   → Works on the ORIGINAL struct
//
// 🧠 RULE OF THUMB:
// Use POINTER receiver when you want to modify the struct’s data.
//

// ✅ changeStatus method using POINTER RECEIVER
func (o *Order) changeStatus(newStatus string) {
	// Here `o` is a pointer to an `Order`.
	// So this line modifies the *actual* struct, not a copy.
	o.status = newStatus

	// 🔍 Note:
	// You could write `(*o).status = newStatus` (manual dereferencing),
	// but Go automatically dereferences pointers when using dot `.`.
	//
	// i.e., o.status = newStatus ✅  ==  (*o).status = newStatus ✅
}

// 🧾 A read-only method with VALUE receiver (for printing details)
func (o Order) printSummary() {
	// Since (o Order) is a VALUE receiver, it works on a COPY of the struct.
	// This is fine because we are only *reading*, not modifying.
	fmt.Printf("Order #%d | Amount: %.2f | Status: %s | Created: %s\n",
		o.id, o.amount, o.status, o.created.Format("2006-01-02 15:04:05"))
}

//////////////////////////////////////////////////////////////////////
// 🧩 STEP 3: MAIN FUNCTION — Demonstrate Everything
//////////////////////////////////////////////////////////////////////

func main() {

	//------------------------------------------------------------
	// 🟢 Way 1: Using shorthand ( := ) → most common way (inside functions)
	//------------------------------------------------------------
	order1 := Order{
		id:      101,
		amount:  2500.75,
		status:  "Pending",
		created: time.Now(),
	}

	fmt.Println("Order 1 (shorthand):", order1)
	fmt.Printf("Formatted output (with field names): %+v\n", order1)
	fmt.Println("Accessing fields individually →", order1.id, order1.amount, order1.status, order1.created)

	//------------------------------------------------------------
	// 🟢 Way 2: Using var keyword (declare empty, assign later)
	//------------------------------------------------------------
	var order2 Order // Zero values: int→0, float64→0.0, string→"", time→0001-01-01...

	// Assigning values one by one
	order2.id = 102
	order2.amount = 1800.50
	order2.status = "Completed"
	order2.created = time.Now()

	fmt.Println("Order 2 (var empty, assign later):", order2)

	//------------------------------------------------------------
	// 🟢 Way 3: Using var + initialization together
	//------------------------------------------------------------
	var order3 = Order{
		id:      103,
		amount:  3000.00,
		status:  "Failed",
		created: time.Now(),
	}
	fmt.Println("Order 3 (var + initialized):", order3)

	//------------------------------------------------------------
	// 🟢 Way 4: Using pointer to struct
	//------------------------------------------------------------
	order4 := &Order{
		id:      104,
		amount:  999.99,
		status:  "Processing",
		created: time.Now(),
	}

	// Even though order4 is a pointer, we can use '.' (dot) directly in Go.
	// Go automatically dereferences the pointer to access its fields.
	fmt.Println("Order 4 (pointer):", *order4)
	fmt.Println("Access via pointer:", order4.id, order4.amount, order4.status, order4.created)

	//------------------------------------------------------------
	// 🧩 BONUS: Updating values via pointer receiver
	//------------------------------------------------------------
	// Let's use our method that has a POINTER receiver.
	order4.changeStatus("Delivered") // ✅ modifies the original struct

	fmt.Println("Order 4 after status change (pointer receiver):", *order4)

	//------------------------------------------------------------
	// 🧾 BONUS: Using value receiver (read-only)
	//------------------------------------------------------------
	order4.printSummary()

	//------------------------------------------------------------
	// 🧩 Understanding Dereferencing in Receiver
	//------------------------------------------------------------
	//
	// Q: Can we manually dereference in receiver? ✅ YES
	// But we don't need to — Go automatically handles it.
	//
	// Example inside a method:
	//     (*o).status = "Delivered"   // manual
	//     o.status = "Delivered"      // automatic (Go does this)
	//
	// Both are valid and equivalent.
	//
	// Similarly, when calling the method:
	//     order4.changeStatus("Done")  // order4 is pointer → direct call
	//     order1.changeStatus("Done")  // order1 is value → Go auto takes &order1
	//
	// ✅ Go auto converts between value ↔ pointer where safe.
	//
	// This keeps syntax clean while preserving performance.
	//
}

// imp//-----------
// In Go, two structs must have the same type (not just same fields) to be directly assignable.

// 🔍 Explanation:

// Even if two structs have identical fields, if their type names are different, they are considered different types — so you can’t assign one to another directly.
