package main

import "fmt"

//////////////////////////////////////////////////////////////////////
// 🧩 Step 1: Define a basic interface — PaymentGateway
//////////////////////////////////////////////////////////////////////
//
// 👉 This interface defines a basic payment behavior — MakePayment().
// 👉 Any type (like PayPal, Stripe, Razorpay) that implements this
//    method will satisfy the PaymentGateway interface.
//
type PaymentGateway interface {
	MakePayment(amount float64)
}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 2: Define another small interface — Refundable
//////////////////////////////////////////////////////////////////////
//
// 👉 This interface defines another behavior — Refund().
// 👉 We can use this for gateways that support refunds.
//
type Refundable interface {
	Refund(amount float64)
}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 3: Compose interfaces — AdvancedPaymentGateway
//////////////////////////////////////////////////////////////////////
//
// 👉 This new interface combines PaymentGateway + Refundable.
// 👉 Any struct that implements both MakePayment() and Refund()
//    will automatically satisfy AdvancedPaymentGateway.
//
type AdvancedPaymentGateway interface {
	PaymentGateway
	Refundable
}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 4: Define concrete types — PayPal, Stripe, Razorpay
//////////////////////////////////////////////////////////////////////
//
// 👉 Each gateway implements its own logic for MakePayment().
// 👉 Only some gateways (like PayPal, Stripe) may support refunds.
//
type PayPal struct{}

func (p PayPal) MakePayment(amount float64) {
	fmt.Printf("💰 Payment of ₹%.2f made via PayPal.\n", amount)
}

func (p PayPal) Refund(amount float64) {
	fmt.Printf("↩️ Refund of ₹%.2f processed via PayPal.\n", amount)
}

type Stripe struct{}

func (s Stripe) MakePayment(amount float64) {
	fmt.Printf("💳 Payment of ₹%.2f made via Stripe.\n", amount)
}

func (s Stripe) Refund(amount float64) {
	fmt.Printf("💵 Refund of ₹%.2f processed via Stripe.\n", amount)
}

type Razorpay struct{}

func (r Razorpay) MakePayment(amount float64) {
	fmt.Printf("🏦 Payment of ₹%.2f made via Razorpay.\n", amount)
}

// ⚠ Razorpay does not implement Refund() method,
// so it only satisfies PaymentGateway, not AdvancedPaymentGateway.

//////////////////////////////////////////////////////////////////////
// 🧩 Step 5: Function that uses interface
//////////////////////////////////////////////////////////////////////
//
// 👉 processPayment() can take *any* type that satisfies PaymentGateway.
// 👉 This means it can handle PayPal, Stripe, or Razorpay.
//
func processPayment(pg PaymentGateway, amount float64) {
	pg.MakePayment(amount)
}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 6: Another function for advanced gateways
//////////////////////////////////////////////////////////////////////
//
// 👉 refundPayment() only accepts AdvancedPaymentGateway types.
// 👉 So Razorpay (which lacks Refund()) can’t be passed here.
//
func refundPayment(pg AdvancedPaymentGateway, amount float64) {
	pg.Refund(amount)
}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 7: Main function — Demonstrate everything
//////////////////////////////////////////////////////////////////////
//
// 👉 We’ll create instances of each payment gateway and show how
//    interface composition helps us write flexible, reusable code.
//
func main() {
	paypal := PayPal{}
	stripe := Stripe{}
	razorpay := Razorpay{}

	fmt.Println("✅ --- Making Payments ---")
	processPayment(paypal, 500.00)
	processPayment(stripe, 1000.00)
	processPayment(razorpay, 750.00)

	fmt.Println("\n↩️ --- Processing Refunds ---")
	refundPayment(paypal, 200.00)
	refundPayment(stripe, 400.00)

	// ⚠ This line would cause an error because Razorpay doesn't implement Refund():
	// refundPayment(razorpay, 100.00)

}

/*
🧠 Ek simple analogy (real life example)

Soch:

PaymentGateway = koi bhi jo payment le sakta hai 💰

AdvancedPaymentGateway = jo payment bhi le aur refund bhi kar sake ↩️

Agar tu ek “basic payment” lena chahta hai,
to tujhe har baar refund karne wale system ki zarurat nahi hoti.

Tu chahe PayPal le, Stripe le, Razorpay le — sab chalega.

Par agar tu kehta hai “Sirf wo gateway chalega jo refund bhi kar sake”,
to Razorpay automatically bahar ho jaata hai.

*/
