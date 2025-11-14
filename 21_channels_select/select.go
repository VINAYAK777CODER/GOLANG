package main

import (
	"fmt"
	"time"
)

func main() {

	// Create two channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	// -------------------------------
	// GOROUTINE 1
	// Sends a value to ch1 AFTER 1 second
	// -------------------------------
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from ch1"
	}()

	// -------------------------------
	// GOROUTINE 2
	// Sends a value to ch2 AFTER 1 second
	// -------------------------------
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from ch2"
	}()

	// --------------------------------------------------------------------
	// SELECT STATEMENT
	// It listens to MULTIPLE channels AT THE SAME TIME.
	// --------------------------------------------------------------------
	// IMPORTANT:
	// select checks all cases and chooses ONE that is ready.
	// --------------------------------------------------------------------
	// ⭐ IMPORTANT RULE ⭐
	// If multiple select cases are ready at the same time,
	// -> Go chooses ONE case RANDOMLY.
	// -> ONLY ONE case executes.
	//
	// This rule applies ONLY when:
	// - Multiple channels already have data ready at the SAME MOMENT.
	//
	// This rule does NOT apply when:
	// - One channel becomes ready earlier than the other
	//   (in that case select chooses the earliest one, no randomness).
	//
	// In our example:
	// Both goroutines sleep for EXACTLY 1 second.
	// So both channels become ready at the "same time".
	// Therefore: select will pick ONE randomly.
	// --------------------------------------------------------------------

	select {
	case msg := <-ch1:
		// This block runs IF ch1's data was selected
		fmt.Println("Received:", msg)

	case msg := <-ch2:
		// This block runs IF ch2's data was selected
		fmt.Println("Received:", msg)
	}

	// NOTE:
	// You will get ONLY ONE output:
	// Either:
	//   Received: Message from ch1
	// OR:
	//   Received: Message from ch2
	//
	// select NEVER executes both cases in one call.
}




/*
package main

import (
	"fmt"
	"time"
)

func main() {

	// Store the time when the program starts.
	// We will use this to print how much time has passed.
	start := time.Now()

	// tick: A channel that sends a signal every 100 milliseconds.
	// It's like a timer that repeatedly "ticks".
	tick := time.Tick(100 * time.Millisecond)

	// boom: A channel that sends ONE signal after 500 milliseconds.
	// After it sends the signal once, it will never send again.
	boom := time.After(500 * time.Millisecond)

	// elapsed(): A small function that returns
	// how much time has passed since the program started.
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}

	// Infinite loop that keeps checking which event happened.
	for {
		select {

		// ------------------------------
		// CASE 1: tick event
		// This runs every time 100ms passes.
		// ------------------------------
		case <-tick:
			fmt.Printf("[%6s] tick.\n", elapsed())

		// ------------------------------
		// CASE 2: boom event
		// This will happen ONCE at exactly 500ms.
		// When this happens, the program prints "BOOM!" and exits.
		// ------------------------------
		case <-boom:
			fmt.Printf("[%6s] BOOM!\n", elapsed())
			return // Exit the program

		// ------------------------------
		// DEFAULT CASE
		// This runs when NEITHER tick nor boom is ready.
		// Meaning: select does NOT block waiting.
		// ------------------------------
		default:
			// Print a dot to show "nothing happened right now".
			fmt.Printf("[%6s]     .\n", elapsed())

			// Sleep 50ms so the program doesn't run too fast.
			time.Sleep(50 * time.Millisecond)
		}
	}
}



*/