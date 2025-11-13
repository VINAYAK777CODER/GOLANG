package main

import "fmt"

////////////////////////////////////////////////////////////////////////////////
// 🧩 Unbuffered Channel
//
// ❌ No storage — can’t hold any data.
// ⏸️ Send blocks until receiver ready.
// ⏸️ Receive blocks until sender sends.
// ⚡ Used for synchronization / handshakes.
// ⚠️ Deadlock if both send & receive in same goroutine.
// ✅ Perfect for signaling completion or waiting.
// 🧠 Think: Phone call — both must be present to talk.
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
// 🧺 Buffered Channel
//
// ✅ Has storage (buffer capacity > 0).
// 🚀 Send doesn’t block until buffer full.
// 🕒 Receive doesn’t block until buffer empty.
// 💡 Used for asynchronous communication (producer–consumer).
// 🧮 Controls throughput by setting buffer size.
// 🧠 Think: Mailbox — drop letters, pick later.
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
// ⚖️ Comparison Table
//
// Feature                | Unbuffered         | Buffered
// --------------------------------------------------------------
// Capacity               | 0                  | >0
// Send Blocks When       | No receiver        | Buffer full
// Receive Blocks When    | No sender          | Buffer empty
// Communication          | Synchronous        | Asynchronous
// Use Case               | Sync / signaling   | Pipelines / queues
// Analogy                | Phone call 🤝      | Mailbox 📬
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
// ⚡ Deadlock Rule
//
// If you use an unbuffered channel without another goroutine to receive → ❌ deadlock!
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
// 🧠 In One Line
//
// 🧩 Unbuffered = handshake (wait each other)
// 🧺 Buffered   = mailbox (store & continue)
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
// ✅ Remember:
//
// - Use unbuffered → for tight synchronization.
// - Use buffered   → for performance & async processing.
// - Always close channels when no more data will be sent.
////////////////////////////////////////////////////////////////////////////////

func main() {

	//////////////////////////////////////////////////////////////////////
	// 🧩 Example 1: Unbuffered Channel
	//////////////////////////////////////////////////////////////////////
	ch1 := make(chan string)

	go func() {
		ch1 <- "Hello from Unbuffered Channel" // 🚀 send (will block until received)
	}()

	fmt.Println(<-ch1) // 📨 receive (unblocks sender)

	//////////////////////////////////////////////////////////////////////
	// 🧺 Example 2: Buffered Channel
	//////////////////////////////////////////////////////////////////////
	ch2 := make(chan string, 2) // buffer capacity = 2

	ch2 <- "Message 1"
	ch2 <- "Message 2"
	// ch2 <- "Message 3" // ❌ would block here (buffer full)

	fmt.Println(<-ch2) // receives first message
	fmt.Println(<-ch2) // receives second message

	fmt.Println("🏁 Program completed successfully!")
}
