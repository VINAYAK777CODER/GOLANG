// 🧩 Topic: Understanding Goroutines, Channels, and Closing Channels in Go
// -------------------------------------------------------
// In this example, we simulate sending emails using a goroutine.
// We'll explain:
//   1️⃣ What the code does
//   2️⃣ The mistake (missing close)
//   3️⃣ Why time.Sleep() doesn’t fix it
//   4️⃣ How to fix it properly
// -------------------------------------------------------

package main

import (
	"fmt"
	"time"
)

// -----------------------------------------------
// 🧠 FUNCTION: emailSender
// -----------------------------------------------
// This function runs as a goroutine.
// It receives email addresses from a channel and simulates sending them.
// The 'done' channel is used to signal the main function that
// all emails have been processed.
//
// PARAMETERS:
//   - email: channel carrying email addresses (string)
//   - done: channel used to signal completion (bool)
// -----------------------------------------------
func emailSender(email chan string, done chan bool) {

	// 🧩 'defer' ensures that when the function exits,
	// it sends a signal 'true' to the 'done' channel.
	// This is how the main goroutine knows we’re finished.
	defer func() {
		done <- true
	}()

	// 🧠 'for e := range email' means:
	// Keep receiving from the 'email' channel until it’s closed.
	// Once closed, the loop automatically ends.
	for e := range email {
		fmt.Println("sending email to:", e)

		// Simulate time taken to send each email.
		time.Sleep(time.Second)
	}

	// When the channel is closed and drained, this loop ends.
	// Then defer sends 'true' → main() knows work is done.
}

// -----------------------------------------------
// 🧩 FUNCTION: main
// -----------------------------------------------
func main() {

	// Create a buffered channel 'email' with capacity 100.
	// 🧠 A buffered channel can hold up to 100 emails without blocking.
	email := make(chan string, 100)

	// 'done' channel (unbuffered) for synchronization.
	done := make(chan bool)

	// Start the goroutine that will send emails.
	// It runs concurrently with the main function.
	go emailSender(email, done)

	// 🧠 This loop sends 101 emails into the channel (0 to 100).
	// Since the buffer is 100, the 101st send will block until
	// the goroutine starts reading from the channel.
	for i := 0; i <= 100; i++ {
		email <- fmt.Sprintf("%d@gmail.com", i)
	}

	// --------------------------------------------------
	// ❌ MISTAKE IN ORIGINAL CODE (if 'close(email)' missing):
	//
	// If we don’t close the channel here:
	//   → 'emailSender' keeps waiting for more emails forever.
	//   → It never exits the 'for range' loop.
	//   → The 'defer done<-true' never executes.
	//   → main() waits forever at '<-done' → DEADLOCK.
	//
	// ❗ Adding time.Sleep() inside 'emailSender' does NOT fix it.
	// It only slows down email sending — it does NOT stop
	// the goroutine from waiting forever after the channel empties.
	// --------------------------------------------------

	// ✅ CORRECT FIX: Close the 'email' channel
	// This tells 'emailSender' that no more data will come.
	close(email)

	// Print confirmation after sending all emails.
	fmt.Println("Done sending")

	// 🧠 This line blocks until 'emailSender' finishes and sends 'true' to 'done'.
	// So main() will not exit until all emails are processed.
	<-done
}

// -----------------------------------------------
// 🧩 FINAL WORKFLOW SUMMARY
// -----------------------------------------------
//
// STEP 1️⃣: main() creates channels: 'email' and 'done'.
//
// STEP 2️⃣: main() starts 'emailSender' as a goroutine.
//
// STEP 3️⃣: main() pushes 101 email addresses into 'email'.
//
// STEP 4️⃣: main() closes 'email'.
//           🔹 This is crucial because 'for e := range email'
//              keeps waiting until the channel is closed.
//
// STEP 5️⃣: 'emailSender' keeps reading from 'email'.
//           For each email, it prints a message and waits 1 second.
//
// STEP 6️⃣: Once all emails are read and the channel is empty + closed,
//           'for range' ends, and the 'defer' runs → done <- true
//
// STEP 7️⃣: main() unblocks at '<-done' and exits gracefully.
//
// ✅ NO DEADLOCK, ✅ CLEAN EXIT, ✅ ALL EMAILS SENT.
//
// -----------------------------------------------
//
// 🧩 Key Lessons:
//
// 🔹 Always close the sending channel when done sending values.
// 🔹 Receivers using 'for range ch' exit automatically when channel is closed.
// 🔹 'time.Sleep' only delays execution, it doesn’t signal completion.
// 🔹 Use 'done' channel (or WaitGroup) for synchronization.
//
// -----------------------------------------------
//
// 🧩 SAMPLE OUTPUT:
//
// sending email to: 0@gmail.com
// sending email to: 1@gmail.com
// ...
// sending email to: 100@gmail.com
// Done sending
//
// (Program exits cleanly)
//
// -----------------------------------------------
