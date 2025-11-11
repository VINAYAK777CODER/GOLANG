package main

import (
	"fmt"
)

//////////////////////////////////////////////////////////////////////
// 🧩 Step 1: Define a function that will run in a separate goroutine
//////////////////////////////////////////////////////////////////////
//
// This function takes a **channel of strings** as a parameter.
// A channel is like a "pipe" that can send data between goroutines.
//
// The function will send a message into the channel.
func sendmessage(ch chan string) {
	// 📨 Send a string value into the channel.
	// This operation is **blocking** — it will pause here
	// until another goroutine is ready to receive the value from the channel.
	ch <- "hello i am vinayak"
}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 2: main() function — program entry point
//////////////////////////////////////////////////////////////////////
func main() {

	// 🏗️ Create a new channel that can carry string values.
	// `make(chan string)` allocates a new **unbuffered channel**.
	// "Unbuffered" means: send and receive must happen at the same time
	// (otherwise the goroutine will wait).
	ch := make(chan string)

	// 🚀 Launch a new goroutine that runs the sendmessage() function.
	// This runs **concurrently** with main().
	// The new goroutine and the main goroutine will now run in parallel.
	go sendmessage(ch)

	// 🕒 Now, main() is waiting to receive a value from the channel.
	// `<-ch` means: "Receive a value from channel ch".
	// This is also a **blocking operation** — it will pause until
	// the sendmessage() goroutine sends something into the channel.
	msg := <-ch

	// 🧾 Once the message is received, the blocking is over.
	// The received string is stored in variable `msg`.

	// 🖨️ Print the received message to the console.
	fmt.Println(msg)
}

//////////////////////////////////////////////////////////////////////
// ⚙️ HOW THE PROGRAM EXECUTES (Step-by-step timeline)
//////////////////////////////////////////////////////////////////////
//
// 1️⃣ main() starts and creates a channel: ch := make(chan string)
//
// 2️⃣ main() launches a goroutine: go sendmessage(ch)
//     → Now, Go runtime starts sendmessage() in a separate thread of execution.
//
// 3️⃣ sendmessage() runs and executes:
//        ch <- "hello i am vinayak"
//     This tries to **send** data into the channel.
//
// 4️⃣ At the same time, main() executes:
//        msg := <-ch
//     This tries to **receive** data from the channel.
//
// 🧩 Since both send and receive are ready, data transfer happens immediately:
//     "hello i am vinayak" moves from sendmessage() → main()
//
// 5️⃣ sendmessage() completes and exits after sending the message.
//
// 6️⃣ main() receives the message, stores it in `msg`, and prints it.
//
// Output:
//     hello i am vinayak
//
// ✅ Synchronization happens automatically using the channel —
//    there is no need for any sleep(), waitgroup, or mutex.
//////////////////////////////////////////////////////////////////////
