package main

import (
	"fmt"
)

//////////////////////////////////////////////////////////////////////
// 🧩 Step 1: Define a function that sends multiple messages
//////////////////////////////////////////////////////////////////////
//
// This function takes a channel as a parameter.
// It will send multiple string messages using a loop.
func sendMessages(ch chan string) {
	// 📦 Send 5 messages one by one through the channel
	for i := 1; i <= 5; i++ {
		// Format the message dynamically
		message := fmt.Sprintf("📨 Message %d from Vinayak", i)

		// Send the message into the channel
		// ⚠️ This will block until the receiver (main goroutine)
		//     reads the value from the channel.
		ch <- message
		fmt.Println("✅ Sent:", message)
	}

	// 🚪 After sending all messages, close the channel.
	// This tells the receiver that no more data will be sent.
	close(ch)
}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 2: main() — the receiver goroutine
//////////////////////////////////////////////////////////////////////
func main() {
	// 🏗️ Create a channel to carry string data
	ch := make(chan string)

	// 🚀 Start the sending goroutine
	go sendMessages(ch)

	// 🕒 Receive messages from the channel inside a loop
	//
	// The `range` keyword automatically receives values
	// until the channel is closed.
	for msg := range ch {
		fmt.Println("💬 Received:", msg)
	}

	// When the channel is closed and empty, `range` stops automatically.
	fmt.Println("✅ All messages received. Channel closed.")
}
/*
⚙️ Step-by-Step Execution Flow

1️⃣ main() creates a channel → ch := make(chan string)
2️⃣ go sendMessages(ch) starts a new goroutine to send messages.
3️⃣ Inside that goroutine, the loop sends 5 messages → ch <- message.
4️⃣ Each send blocks until main() receives it.
5️⃣ main() uses for msg := range ch to continuously receive from the channel.
6️⃣ After all 5 sends, close(ch) is called → tells receiver that sending is done.
7️⃣ The range loop ends automatically when the channel is closed and empty.
*/