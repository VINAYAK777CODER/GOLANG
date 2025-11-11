package main

import (
	"fmt"
	"sync"
	"time"
)

//////////////////////////////////////////////////////////////////////
// 🧩 Step 1: Define the function that will run in a goroutine
//////////////////////////////////////////////////////////////////////
//
// Each goroutine represents one "task" (like an API call or process).
// The WaitGroup keeps track of how many such goroutines are running.
//
// When a goroutine finishes, it must call wg.Done()
// to tell the WaitGroup: “I’m done! Decrease the counter by 1.”
//
func taskid(id int, wg *sync.WaitGroup) {
	// 🔹 defer means this will run automatically when the function ends
	defer wg.Done()

	fmt.Println("🔹 [STARTED] Task:", id)

	// Simulate some work (each goroutine sleeps independently)
	time.Sleep(time.Millisecond * 500)

	fmt.Println("✅ [FINISHED] Task:", id)
}

//////////////////////////////////////////////////////////////////////
// 🧩 Step 2: The main goroutine (program entry point)
//////////////////////////////////////////////////////////////////////
func main() {

	//////////////////////////////////////////////////////////////////////
	// 🧠 STEP 2.1: Create a WaitGroup instance
	//////////////////////////////////////////////////////////////////////
	//
	// WaitGroup acts like a counter.
	// It starts at 0. Each time we say wg.Add(1), we increment it.
	// Each time a goroutine calls wg.Done(), we decrement it.
	// wg.Wait() will block until the counter becomes 0.
	//
	var wg sync.WaitGroup

	//////////////////////////////////////////////////////////////////////
	// 🧠 STEP 2.2: Launch multiple goroutines
	//////////////////////////////////////////////////////////////////////
	//
	// Imagine we have 5 tasks to perform concurrently.
	// Before launching each goroutine, we must increment the counter
	// because we’re saying, “There will be one more goroutine to wait for.”
	//
	for i := 1; i <= 5; i++ {

		// 🔸 Increase the counter by 1.
		// Internally WaitGroup counter = counter + 1
		wg.Add(1)

		// 🔸 Launch goroutine (runs concurrently with others)
		go taskid(i, &wg)

		// ⏱ By now:
		// Iteration 1 → counter = 1
		// Iteration 2 → counter = 2
		// ...
		// Iteration 5 → counter = 5
		//
		// So after the loop, counter = 5 (5 tasks are running)
	}

	//////////////////////////////////////////////////////////////////////
	// 🧠 STEP 2.3: Wait for all goroutines to finish
	//////////////////////////////////////////////////////////////////////
	//
	// The line wg.Wait() tells the main goroutine:
	// “Pause here and don’t continue until the counter becomes 0.”
	//
	// Each goroutine, when it finishes its work, calls wg.Done(),
	// which decreases the counter by 1.
	//
	// So the workflow looks like this:
	//
	// Initial counter = 5
	// Task 1 done → counter = 4
	// Task 2 done → counter = 3
	// Task 3 done → counter = 2
	// Task 4 done → counter = 1
	// Task 5 done → counter = 0  ← when this happens, wg.Wait() unblocks!
	//
	wg.Wait()

	//////////////////////////////////////////////////////////////////////
	// 🧠 STEP 2.4: Once all are done, program continues
	//////////////////////////////////////////////////////////////////////
	//
	// Now all goroutines have finished their jobs.
	// The main goroutine can safely continue and end the program.
	//
	fmt.Println("🎉 All goroutines finished! Main exiting.")
}
