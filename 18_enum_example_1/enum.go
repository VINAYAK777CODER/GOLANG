package main

import "fmt"

// ------------------------------------------------------------
// 🧱 STEP 1: Define a Custom Type
// ------------------------------------------------------------
// We create a new type 'OrderStatus' which is based on 'int'.
// Internally, it will still store numbers (0, 1, 2, 3...),
// but Go will treat 'OrderStatus' as a distinct type — not a plain int.
// This gives us type safety and makes our code more readable.
type OrderStatus int

// ------------------------------------------------------------
// 🧮 STEP 2: Define Enum Values using iota
// ------------------------------------------------------------
// iota starts from 0 and automatically increments by 1 for each line.
// This means:
// Pending   = 0
// Shipped   = 1
// Delivered = 2
// Cancelled = 3
//
// These are constants of type 'OrderStatus' (not plain integers).
const (
	Pending OrderStatus = iota
	Shipped
	Delivered
	Cancelled
)

// ------------------------------------------------------------
// 🧩 STEP 3: Add a String() Method to Convert Int → Readable String
// ------------------------------------------------------------
// Go's fmt package automatically checks if a type implements this method:
//     func (T) String() string
//
// If yes, then when you print the variable (fmt.Println),
// Go calls this function internally to get a readable string.
//
// Example:
// fmt.Println(Pending) → calls Pending.String() → returns "Pending"
func (s OrderStatus) String() string {
	switch s {
	case Pending:
		return "Pending"
	case Shipped:
		return "Shipped"
	case Delivered:
		return "Delivered"
	case Cancelled:
		return "Cancelled"
	default:
		return "Unknown" // handles invalid values safely
	}
}

// ------------------------------------------------------------
// ⚙️ STEP 4: A Function That Uses the Enum
// ------------------------------------------------------------
// This function simulates order progress. Based on the current order status,
// it returns the next logical state.
//
// Note: Even though it returns 'OrderStatus' (which is int underneath),
// when you print it, the String() method will automatically
// convert it into human-readable text.
func getNextStatus(current OrderStatus) OrderStatus {
	switch current {
	case Pending:
		return Shipped   // 0 → 1
	case Shipped:
		return Delivered // 1 → 2
	default:
		return current   // stay the same for other statuses
	}
}

// ------------------------------------------------------------
// 🏁 STEP 5: Main Function — Program Entry Point
// ------------------------------------------------------------
func main() {

	// Step 5.1 — Initialize the order with Pending status.
	// Internally, 'Pending' = 0 (OrderStatus type).
	status := Pending
	fmt.Println("Initial:", status)
	// 👉 Output → "Initial: Pending"
	// Explanation:
	// fmt.Println detects the String() method, calls status.String(),
	// which returns "Pending".

	// Step 5.2 — Move to the next status.
	status = getNextStatus(status)
	fmt.Println("Next:", status)
	// 👉 Output → "Next: Shipped"
	// Internally:
	// getNextStatus(Pending) returned 'Shipped' (1).
	// fmt.Println calls Shipped.String() → "Shipped".

	// Step 5.3 — Move again to the next status.
	status = getNextStatus(status)
	fmt.Println("Next:", status)
	// 👉 Output → "Next: Delivered"
	// Internally:
	// getNextStatus(Shipped) returned 'Delivered' (2).
	// fmt.Println calls Delivered.String() → "Delivered".
}

/*
===============================================================
🧠  SUMMARY OF CODE FLOW
===============================================================
1️⃣  type OrderStatus int
     → Creates a custom enum type stored as int internally.

2️⃣  const (...) = iota
     → Assigns auto-incrementing numeric constants:
        Pending=0, Shipped=1, Delivered=2, Cancelled=3.

3️⃣  String() method
     → Translates those numeric values into readable words.
     → fmt.Println() automatically uses this method.

4️⃣  getNextStatus()
     → Returns the next enum (numeric) value.

5️⃣  main()
     → Demonstrates how enum changes and prints readable status.

===============================================================
🧩 INTERNAL BEHAVIOR vs PRINTED OUTPUT
===============================================================
| Enum Name | Internal Value | What You See When Printed |
|------------|----------------|----------------------------|
| Pending    | 0              | "Pending"                 |
| Shipped    | 1              | "Shipped"                 |
| Delivered  | 2              | "Delivered"               |
| Cancelled  | 3              | "Cancelled"               |

✅ In memory → numbers (int)
✅ On screen → readable strings (via String() method)
===============================================================
*/
