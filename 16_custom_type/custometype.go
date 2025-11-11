// main.go
package main

import (
	"fmt"
)

/*
===============================================================================
TL;DR (cheat-sheet inside code)
// -> Ye section sirf comments hai, compile par koi asar nahi 🙂

- type Status string  → naya **defined type** banata hai (base: string, identity: alag).
- Is par tum **methods** laga sakte ho:  func (s Status) IsActive() bool
- Status **string ke barabar nahi** hota (explicit conversion chahiye).
- Best use: readability, type-safety, enum-like constants.

1) Keywords/Terms (seedha simple)
- Defined type:       type Status string      → Status ek naya type, base: string
- Alias:              type Status = string    → sirf dusra naam, naya type nahi
- Receiver method:    func (s Status) ...     → method Status ke saath attach

2) Sample program ka stepwise flow (jo niche main() me run bhi hoga)
- type Status string
  -> Compiler ko bolo: “Status naam ka naya type, jo andar se string jaisa behave kare.”
- func (s Status) IsActive() bool { return s == "Active" }
  -> Har Status value ke paas ab IsActive() method hoga.
  -> Call style: order.IsActive(); internally: IsActive(order).
- var order Status = "Active"
  fmt.Println(order.IsActive())  // true
  Flow:
    1) order ban gaya (Status type) with value "Active"
    2) order.IsActive() call hua
    3) Function check: order == "Active" → true
    4) true print ho gaya
  Agar order = "Inactive" karoge, IsActive() → false

3) Kyun use karein? (real benefits)
- Type-safety: Status ≠ string → galat assignment se bachega.
- Readability: Status naam se intent clear hota hai (random string nahi).
- Methods: domain logic type ke saath chipka sakte ho (IsActive, CanCancel, etc.).
- Enum-like constants: fixed valid values define kar sakte ho.

4) Defined type vs Alias (most asked)
+----------------------+-------------------+------------------------+---------------------------+
| Syntax               | Meaning           | Equal to base type?    | Methods attach possible?  |
+----------------------+-------------------+------------------------+---------------------------+
| type Status string   | New defined type  | ❌ (conversion needed) | ✅ (methods allowed)       |
| type Status = string | Alias (same type) | ✅ (same as string)    | ❌ (alias par methods nahi)|
+----------------------+-------------------+------------------------+---------------------------+

Conversion example (defined type):
    var s Status = "Active"
    var t string
    // t = s      // ❌ compile error
    t = string(s) // ✅ explicit conversion

Extra tips:
- Value receiver vs pointer receiver:
  - Value:   read-only style (func (s Status) IsActive() bool)
  - Pointer: modify karna ho (func (s *Status) Activate())
- Always prefer constants for allowed values → typos se bachoge.
===============================================================================
*/

// ----------------------------------------------------------------------------
//  A) Defined type: Status (base type = string)
// ----------------------------------------------------------------------------
type Status string

// Enum-like constants (fixed valid values)
// 👉 Best practice: string literals ke bajay in constants ko hi use karo.
const (
	Active   Status = "Active"
	Inactive Status = "Inactive"
	Pending  Status = "Pending"
)

// Value receiver method: sirf read/check karna ho to value receiver theek.
// Yahan pe hum check kar rahe hain ki kya Status == "Active".
func (s Status) IsActive() bool {
	// NOTE: Direct literal "Active" ki jagah constant Active use karna aur bhi safe hai:
	// return s == Active
	return s == "Active"
}

// Extra read-only helpers (clean, readable API)
func (s Status) IsPending() bool  { return s == Pending }
func (s Status) IsInactive() bool { return s == Inactive }

// Pointer receiver method: modify karna ho to pointer use karo.
func (s *Status) Activate()   { *s = Active }
func (s *Status) Deactivate() { *s = Inactive }
func (s *Status) MarkPending() { *s = Pending }

// ----------------------------------------------------------------------------
//  B) Alias type vs Defined type (IMPORTANT DIFFERENCE)
// ----------------------------------------------------------------------------

// ❗ Alias declaration: yeh "naya type" nahi banata, bas dusra naam deta hai.
// Is wajah se alias par methods ATTACH nahi kar sakte.
// Try karoge to compile error aayega agar method receiver alias ho.
type StatusAlias = string // <- Alias: StatusAlias bilkul string ke barabar hoga

// Uncomment karke dekho (will NOT compile):
// func (a StatusAlias) Nope() {} // ❌ methods cannot be defined on alias of a non-local type

// ----------------------------------------------------------------------------
//  C) Small demo helpers for printing
// ----------------------------------------------------------------------------
func printHeader(title string) {
	fmt.Println("\n============================================================")
	fmt.Println(title)
	fmt.Println("============================================================")
}

func main() {
	// ------------------------------------------------------------------------
	// 1) Stepwise flow demo (exactly wahi jo notes me likha tha)
	// ------------------------------------------------------------------------
	printHeader("1) Stepwise flow demo")

	// Step 1: order ban gaya (Status type) with value "Active"
	var order Status = "Active" // ya var order Status = Active (recommended constants)
	fmt.Println("order (value):", order)

	// Step 2: order.IsActive() call hua (internally IsActive(order))
	res := order.IsActive()

	// Step 3: function check karta hai order == "Active"
	// Step 4: true print hota hai
	fmt.Println("order.IsActive():", res) // expect: true

	// Ab value change karke dekho
	order = "Inactive" // ya Inactive
	fmt.Println("order set to:", order)
	fmt.Println("order.IsActive():", order.IsActive()) // expect: false

	// ------------------------------------------------------------------------
	// 2) Type-safety demo (Status ≠ string)
	// ------------------------------------------------------------------------
	printHeader("2) Type-safety demo (Status ≠ string)")

	var s Status = Active  // Status type
	var t string = "Active" // plain string type

	fmt.Println("s (Status):", s)
	fmt.Println("t (string):", t)

	// ❌ Direct assign nahi chalega: string <- Status (different types)
	// t = s // <-- is line ko uncomment karoge to compile error aayega:
	// cannot use s (variable of type Status) as type string in assignment

	// ✅ Explicit conversion required:
	t = string(s)
	fmt.Println("t after explicit conversion from Status:", t)

	// Aur ulta bhi:
	// var s2 Status = t // ❌ direct error
	var s2 Status = Status(t) // ✅ explicit conversion to Status
	fmt.Println("s2 (converted from string to Status):", s2)

	// ------------------------------------------------------------------------
	// 3) Clean API with constants + methods (readability showcase)
	// ------------------------------------------------------------------------
	printHeader("3) Clean API with constants + methods")

	var st Status = Pending
	fmt.Println("Initial:", st, "IsPending?", st.IsPending())

	st.Activate() // pointer receiver modifies the value
	fmt.Println("After Activate():", st, "IsActive?", st.IsActive())

	st.Deactivate()
	fmt.Println("After Deactivate():", st, "IsInactive?", st.IsInactive())

	// ------------------------------------------------------------------------
	// 4) Alias vs Defined type quick proof
	// ------------------------------------------------------------------------
	printHeader("4) Alias vs Defined type quick proof")

	// Alias behaves exactly like string
	var a StatusAlias = "I am just a string alias"
	var base string = a // ✅ allowed directly (no conversion)
	fmt.Println("Alias value:", a, "| Assigned to string directly:", base)

	// BUT: alias par methods define nahi kar sakte (code top par commented hai).

	// ------------------------------------------------------------------------
	// 5) Bonus: Function that accepts only Status (not any random string)
	// ------------------------------------------------------------------------
	printHeader("5) Function accepting only Status (type-safety in APIs)")

	printStatus := func(s Status) {
		// API ko pata hai ki Status aayega -> accidental random strings se bach gaye.
		switch s {
		case Active:
			fmt.Println("Order Status:", s, "→ can ship immediately.")
		case Pending:
			fmt.Println("Order Status:", s, "→ waiting for confirmation.")
		case Inactive:
			fmt.Println("Order Status:", s, "→ disabled / archived.")
		default:
			fmt.Println("Order Status:", s, "→ unknown value! avoid literals, use constants.")
		}
	}

	printStatus(Active)
	printStatus(Pending)

	// printStatus("Active") // ❌ compile error: argument type must be Status, not string
	printStatus(Status("RandomLiteral")) // ⚠️ allowed after explicit cast, but discouraged; use constants.

	// ------------------------------------------------------------------------
	// 6) Quick MCQ-style self-check (answers in comments)
	// ------------------------------------------------------------------------
	printHeader("6) MCQ-style self-check")

	fmt.Println("Q1: 'type Status string' ka matlab?")
	fmt.Println("A1: Naya defined type bana jo base 'string' pe hai (methods attach ho sakte).")

	fmt.Println("Q2: Kya Status == string?")
	fmt.Println("A2: Nahi. Explicit conversion chahiye (Status <-> string).")

	fmt.Println("Q3: Alias 'type X = string' par methods attach ho sakte?")
	fmt.Println("A3: Nahi. Alias naya type nahi banata; methods defined type par hi attach hote.")

	fmt.Println("Q4: Value vs Pointer receiver kab?")
	fmt.Println("A4: Value: read-only checks. Pointer: jab modify karna ho (e.g., Activate()).")

	fmt.Println("\nAll done ✅ — ab tum 'defined type' vs 'alias', methods, conversions,")
	fmt.Println("aur enum-like constants ka clear mental model bana chuke ho.")
}
