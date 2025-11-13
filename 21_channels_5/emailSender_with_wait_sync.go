// 🧩 Topic: Receiver-only Parameter in Goroutines (Concurrency in Go)
// -----------------------------------------------------------------------------
// In this program, we simulate an email sending system using multiple workers.
//
// 🔹 main() acts as the SENDER — it sends email addresses into a channel.
// 🔹 emailSender() goroutines act as RECEIVERS — they take email addresses out of the channel and process them.
//
// Here, we use `<-chan string` in the function parameter to make `emailSender()`
// a **receiver-only** goroutine. This ensures that it can only *receive* from
// the channel, not *send* into it.
//
// This makes the function safer, clearer, and more idiomatic in Go.
// -----------------------------------------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

// -----------------------------------------------------------------------------
// 🧠 FUNCTION: emailSender
// -----------------------------------------------------------------------------
// PARAMETERS:
//   id         → worker ID (used for identifying which goroutine is working)
//   emailChan  → `<-chan string` → a **receive-only** channel of type string
//   wg         → pointer to sync.WaitGroup, used to signal when the worker finishes
//
// EXPLANATION OF `<-chan string`:
//   - This means the parameter `emailChan` is *receive-only*.
//   - Inside this function, we can do:
//         e := <-emailChan   ✅ receive value
//     But we CANNOT do:
//         emailChan <- "test@gmail.com" ❌ (compiler error)
//   - This ensures the worker only consumes data, not produces it.
//
// PURPOSE:
//   Each worker receives emails from the shared channel and simulates sending them.
// -----------------------------------------------------------------------------
func emailSender(id int, emailChan <-chan string, wg *sync.WaitGroup) {

	// 🧩 'defer wg.Done()' ensures that when this goroutine finishes
	// (after the channel is closed and all data read),
	// it tells the WaitGroup that it’s done.
	defer wg.Done()

	// 🧠 The 'for e := range emailChan' loop automatically keeps
	// receiving values from the channel until it is closed by the sender.
	// Once closed and drained, this loop ends gracefully.
	for e := range emailChan {

		// Each goroutine (worker) receives *different* emails.
		// The Go runtime automatically distributes messages among workers.
		fmt.Printf("📩 Worker %d sending email to: %s\n", id, e)

		// Simulate sending time — pretend it takes 1 second per email.
		time.Sleep(time.Second)
	}

	// When the channel is closed and empty, the 'for range' exits.
	// The worker then hits the deferred wg.Done() call and signals completion.
}

// -----------------------------------------------------------------------------
// 🧩 FUNCTION: main
// -----------------------------------------------------------------------------
func main() {

	// 🧩 Create a buffered channel with capacity = 100
	// The channel is bidirectional (can send and receive) inside main.
	// But when passed to 'emailSender', it is treated as receive-only.
	emailChan := make(chan string, 100)

	// WaitGroup allows the main goroutine to wait for all workers to finish.
	var wg sync.WaitGroup

	// -------------------------------------------------------------------------
	// 🧩 STEP 1: Launch multiple receiver goroutines (workers)
	// -------------------------------------------------------------------------
	// Here we start 5 concurrent workers.
	// Each one calls 'emailSender' with a receive-only view of 'emailChan'.
	//
	// Since channels in Go are safe for concurrent use, all 5 goroutines
	// will compete to receive from the same 'emailChan'.
	// This is called the "fan-out" pattern — work is spread across workers.
	for i := 1; i <= 5; i++ {
		wg.Add(1)                   // Increase WaitGroup counter by 1
		go emailSender(i, emailChan, &wg) // Launch worker goroutine
	}

	// -------------------------------------------------------------------------
	// 🧩 STEP 2: main() acts as the SENDER
	// -------------------------------------------------------------------------
	// Only main() sends values into the channel.
	// Workers are receiver-only, so they cannot send — compiler ensures safety.
	for i := 0; i < 25; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	// -------------------------------------------------------------------------
	// 🧩 STEP 3: Close the channel
	// -------------------------------------------------------------------------
	// This is very important.
	// Closing tells all receivers: “no more values will be sent.”
	// After this, the 'for e := range emailChan' loops in workers will stop
	// once all emails in the buffer are consumed.
	close(emailChan)

	// -------------------------------------------------------------------------
	// 🧩 STEP 4: Wait for all goroutines to finish
	// -------------------------------------------------------------------------
	// Wait until all 5 workers finish processing their share of emails.
	wg.Wait()

	// After WaitGroup counter = 0 → program proceeds.
	fmt.Println("✅ All emails sent successfully!")
}

// -----------------------------------------------------------------------------
// 🧩 WORKFLOW SUMMARY
// -----------------------------------------------------------------------------
//
// 1️⃣ main() creates a bidirectional channel `emailChan`.
// 2️⃣ main() spawns 5 worker goroutines, each using `<-chan string`
//     → meaning they can *only receive* emails.
// 3️⃣ main() sends 25 email addresses into the channel.
// 4️⃣ Workers read emails concurrently, process them, and simulate sending.
// 5️⃣ main() closes the channel after sending all data.
// 6️⃣ Each worker’s `for range` loop ends when the channel closes.
// 7️⃣ Each worker calls `wg.Done()` (deferred).
// 8️⃣ main() waits (`wg.Wait()`) until all workers are done.
// 9️⃣ Program
