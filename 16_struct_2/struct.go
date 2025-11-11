package main

import (
	"fmt"
	"time"
)

//////////////////////////////////////////////////////////////////////
// 🧩 STEP 1: Define the Struct Type
//////////////////////////////////////////////////////////////////////
//
// 👉 A struct groups multiple fields under one name.
// 👉 Think of it like a lightweight object (without classes).
//
type Order struct {
	id      int
	amount  float64
	status  string
	created time.Time
}

//////////////////////////////////////////////////////////////////////
// 🧩 STEP 2: Constructor Function (Factory Function)
//////////////////////////////////////////////////////////////////////
//
// 🔹 Go doesn't have real constructors like Java/C++,
//    but we can make a **factory function** that returns a new struct.
//
// 🔹 The convention is to name it like `New<TypeName>()`
//
// ✅ This helps when you want to initialize fields properly each time.
//
func NewOrder(id int, amount float64, status string) *Order {
	// Here we return a pointer to a newly created Order struct.
	// So we can modify it easily later.
	order:= Order{
		id:      id,
		amount:  amount,
		status:  status,
		created: time.Now(),
	}
	return &order
}

//////////////////////////////////////////////////////////////////////
// 🧩 STEP 3: Add a Method with Receiver (Pointer Receiver)
//////////////////////////////////////////////////////////////////////
//
// ✅ Pointer receiver is used so that method can modify the original struct.
//
func (o *Order) changeStatus(newStatus string) {
	o.status = newStatus // Go auto-dereferences, so (*o).status = newStatus also works.
}

//////////////////////////////////////////////////////////////////////
// 🧩 STEP 4: Add a Read-only Method with Value Receiver
//////////////////////////////////////////////////////////////////////
func (o Order) printSummary() {
	fmt.Printf("Order #%d | Amount: %.2f | Status: %s | Created: %s\n",
		o.id, o.amount, o.status, o.created.Format("2006-01-02 15:04:05"))
}

//////////////////////////////////////////////////////////////////////
// 🧩 STEP 5: MAIN — Demonstrate Constructor + Inline Struct
//////////////////////////////////////////////////////////////////////
func main() {

	//------------------------------------------------------------
	// ✅ Using Constructor Function (Best Practice)
	//------------------------------------------------------------
	order1 := NewOrder(101, 2500.75, "Pending") // returns *Order (pointer)

	fmt.Println("Using Constructor (NewOrder):")
	order1.printSummary()

	// Change status using pointer receiver method
	order1.changeStatus("Delivered")
	fmt.Println("After status update (pointer receiver):")
	order1.printSummary()

	//------------------------------------------------------------
	// ✅ Inline Struct (Anonymous Struct)
	//------------------------------------------------------------
	//
	// 🔹 Used when you need a quick, one-time struct
	// 🔹 No need to formally declare a new type
	// 🔹 Common in temporary JSON or API data handling
	//
	inlineOrder := struct {
		id      int
		product string
		quantity int
	}{
		id:       5001,
		product:  "Laptop",
		quantity: 2,
	}

	fmt.Println("\nUsing Inline Struct (Anonymous Struct):")
	fmt.Printf("Inline Order → ID: %d | Product: %s | Quantity: %d\n",
		inlineOrder.id, inlineOrder.product, inlineOrder.quantity)

	//------------------------------------------------------------
	// ✅ Bonus: Using Inline Struct inside an existing struct field
	//------------------------------------------------------------
	//
	// You can also embed an inline struct inside another struct.
	//
	type Customer struct {
		name   string
		email  string
		// Inline struct as a field:
		order struct {
			id     int
			amount float64
		}
	}

	customer1 := Customer{
		name:  "Rahul",
		email: "rahul@example.com",
		order: struct {
			id     int
			amount float64
		}{
			id:     201,
			amount: 4500.99,
		},
	}

	fmt.Println("\nCustomer with Inline Struct (Nested):")
	fmt.Printf("Customer: %s (%s)\nOrder ID: %d | Amount: %.2f\n",
		customer1.name, customer1.email, customer1.order.id, customer1.order.amount)

	//------------------------------------------------------------
	// ✅ Summary
	//------------------------------------------------------------
	//
	// Constructor → reusable and clean way to create structs
	// Inline Struct → quick, temporary, no type definition
	//
}
