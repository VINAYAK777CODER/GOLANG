package main

import (
	"fmt"
	"time"
)

//////////////////////////////////////////////////////////////////////
// 🧩 Step 1: Define a normal function that will run as a goroutine
//////////////////////////////////////////////////////////////////////
//
// This function simply prints the ID it receives.
// We'll run multiple copies of this function concurrently using goroutines.
//
func taskid(id int) {
	fmt.Println("Task ID:", id)
}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 2: The main function (entry point)
//////////////////////////////////////////////////////////////////////
//
// By default, Go starts executing from the main() function.
// main() itself runs as the *main goroutine*.
// All other goroutines you start run *concurrently* with it.
//
func main() {

	//////////////////////////////////////////////////////////////////////
	// 🧩 Step 3: Launch multiple goroutines
	//////////////////////////////////////////////////////////////////////
	//
	// The keyword `go` before a function call launches that function
	// as a *goroutine* — meaning it runs independently (in the background).
	//
	// Here, we launch 10 goroutines — each printing a different ID.
	//
	for i := 0; i < 10; i++ {
		go taskid(i) // starts a new goroutine for each iteration
	}

	//////////////////////////////////////////////////////////////////////
	// ⚠️ Step 4: Why Sleep is needed
	//////////////////////////////////////////////////////////////////////
	//
	// The main goroutine (this one) will reach the end of main()
	// immediately after launching those 10 background goroutines.
	//
	// If main() exits, the program terminates instantly —
	// even if the background goroutines haven’t printed yet!
	//
	// To prevent that, we pause the main goroutine briefly
	// using time.Sleep(), giving other goroutines time to finish.
	//
	time.Sleep(time.Second) // waits 1 second before program exits

	//////////////////////////////////////////////////////////////////////
	// 🧩 Step 5: Program ends
	//////////////////////////////////////////////////////////////////////
	//
	// Once the main goroutine finishes sleeping, it exits.
	// The program stops, and any unfinished goroutines are killed.
	//
	// (In real applications, we’d use sync.WaitGroup instead of Sleep.)
	//
}
