package main

import (
	"fmt"
	"sync"
	"time"
)

//////////////////////////////////////////////////////////////////////
// 🧩 Concept: “Concurrent Worker System”
//
// Imagine we’re building a Job Processing System (like what real servers do).
//
// 🎯 What Are We Building?
// --------------------------------------------------
// ✅ 5 jobs (tasks) → numbers 1, 2, 3, 4, 5
// ✅ 3 workers → each worker is a goroutine
// ✅ A channel (jobs) → sends jobs to workers
// ✅ A channel (results) → collects results from workers
// ✅ A WaitGroup → keeps count of active workers and ensures main() waits
//////////////////////////////////////////////////////////////////////

// 🧩 Worker Function
//
// Each worker continuously takes jobs from the 'jobs' channel,
// processes them (simulated with time.Sleep), and sends results into 'results'.
//
// The WaitGroup ensures that main() knows when all workers are done.
func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {

	// When this worker is completely done (no more jobs), mark it as finished.
	defer wg.Done()

	// Keep receiving jobs until the 'jobs' channel is closed and empty.
	for job := range jobs {

		// Worker picks up a job
		fmt.Printf("👷 Worker %d started job %d\n", id, job)

		// Simulate the worker doing time-consuming work (like processing a file or API call)
		time.Sleep(time.Second)

		// Send the result (a message string) back through the results channel.
		results <- fmt.Sprintf("✅ Worker %d finished job %d", id, job)
	}
}

//////////////////////////////////////////////////////////////////////
// 🧩 Main Function
//////////////////////////////////////////////////////////////////////
func main() {

	//////////////////////////////////////////////////////////////////////
	// 🧩 PHASE 1 — Setup (main goroutine only)
	//////////////////////////////////////////////////////////////////////

	// 1️⃣ Create a channel to send jobs to workers.
	// Think of it as a "queue" or "conveyor belt".
	// The buffer size (10) means it can hold up to 10 jobs temporarily.
	jobs := make(chan int, 10)

	// 2️⃣ Create a channel to collect processed job results.
	results := make(chan string)

	// 3️⃣ Create a WaitGroup to synchronize (wait for) all worker goroutines.
	var wg sync.WaitGroup

	//////////////////////////////////////////////////////////////////////
	// 🧩 PHASE 2 — Spawning the Workers
	//////////////////////////////////////////////////////////////////////

	// Start 3 worker goroutines.
	for w := 1; w <= 3; w++ {

		// Tell the WaitGroup that one more worker is active.
		wg.Add(1)

		// Launch the worker goroutine.
		go worker(w, jobs, results, &wg)
	}

	//////////////////////////////////////////////////////////////////////
	// 🧩 PHASE 3 — Sending the Jobs
	//////////////////////////////////////////////////////////////////////

	// Send 5 jobs (numbers 1 to 5) to the 'jobs' channel.
	for j := 1; j <= 5; j++ {
		jobs <- j // 📦 Place job into the channel
		fmt.Println("📨 Sent job", j)
	}

	// Close the 'jobs' channel to signal that no new jobs will be sent.
	close(jobs)

	//////////////////////////////////////////////////////////////////////
	// 🧩 PHASE 4 — WaitGroup + Results Channel Coordination
	//////////////////////////////////////////////////////////////////////

	// Start a separate goroutine that waits for all workers to finish.
	// Once they are all done, it will close the results channel.
	go func() {
		wg.Wait()      // ⏳ Wait until all workers call wg.Done()
		close(results) // 🚪 Close the results channel (no more data will come)
	}()

	//////////////////////////////////////////////////////////////////////
	// 🧩 PHASE 5 — Receiving and Printing Results
	//////////////////////////////////////////////////////////////////////

	// Main goroutine continuously reads from 'results' channel.
	// 'range results' will automatically stop when 'results' is closed.
	for res := range results {
		fmt.Println(res)
	}

	//////////////////////////////////////////////////////////////////////
	// 🏁 Final Phase — Program Completion
	//////////////////////////////////////////////////////////////////////
	fmt.Println("🏁 All jobs processed successfully!")
}

//////////////////////////////////////////////////////////////////////
// ⚙️ DETAILED WORKFLOW (EXACTLY WHAT HAPPENS)
//////////////////////////////////////////////////////////////////////

// 🧩 PHASE 1 — SETUP
// ----------------------------------------------------
// main() creates:
//   → jobs channel (for sending tasks)
//   → results channel (for collecting output)
//   → WaitGroup (to track workers)
//
// At this moment, no goroutines are running yet.


// 🧩 PHASE 2 — SPAWNING WORKERS
// ----------------------------------------------------
// Loop runs 3 times → creates 3 workers.
// Each worker runs 'worker()' function in a separate goroutine.
// Each worker executes:
//     for job := range jobs { ... }
// → which means “I’ll wait here until I receive a job from the 'jobs' channel.”
//
// ✅ Workers are now waiting and idle (blocked on channel read).


// 🧩 PHASE 3 — SENDING JOBS
// ----------------------------------------------------
// main() starts sending jobs 1 → 5 into 'jobs' channel.
// Each 'jobs <- j' operation sends a value into the channel.
//
// ⚙️ Timeline Visualization:
//
// Time   | Action                        | Explanation
// -------|--------------------------------|----------------------------------
// t=0s   | main sends job 1              | worker #1 picks it immediately
// t=0s   | main sends job 2              | worker #2 picks it
// t=0s   | main sends job 3              | worker #3 picks it
// t=0s   | main sends job 4              | job 4 waits in queue
// t=0s   | main sends job 5              | job 5 waits in queue
//
// ✅ At this moment:
// - Worker 1 → busy with job 1
// - Worker 2 → busy with job 2
// - Worker 3 → busy with job 3
// - Job 4, 5 are waiting in the 'jobs' channel buffer.
// - main() closes the jobs channel → means "no new jobs will come."



// 🧩 PHASE 4 — WORKERS PROCESS JOBS
// ----------------------------------------------------
// All workers run concurrently:
//
// Worker 1 → does job 1 → takes 1s → sends "done" message → takes next job (job 4)
// Worker 2 → does job 2 → takes 1s → sends "done" message → takes next job (job 5)
// Worker 3 → does job 3 → takes 1s → sends "done" message → finds no job left → exits
//
// As soon as a job finishes, the result is pushed into 'results' channel.
// main() (below) receives from 'results' and prints it.
//
// After a worker finishes all jobs and the 'jobs' channel is empty,
// its 'for job := range jobs' loop ends → defer wg.Done() executes → reduces WaitGroup count by 1.


// 🧩 PHASE 5 — WAITGROUP & RESULTS CLOSURE
// ----------------------------------------------------
// Meanwhile, another goroutine is waiting inside:
//     wg.Wait()
// Once all 3 workers have called wg.Done() → WaitGroup counter = 0
// It then executes: close(results)
//
// Closing the 'results' channel signals main() that there will be no more outputs.


// 🧩 PHASE 6 — RECEIVING RESULTS
// ----------------------------------------------------
// main() is running:
//     for res := range results
//
// Each time a worker sends a result → main prints it.
// When 'results' is closed and empty → the loop exits automatically.
//
// Finally, we print:
//     🏁 All jobs processed successfully!


// 🧠 VISUAL SUMMARY (like a movie)
//
// main() creates channels + WaitGroup
// |
// |--> starts 3 workers (they wait for jobs)
// |
// |--> sends job 1  → worker 1 picks it up
// |--> sends job 2  → worker 2 picks it up
// |--> sends job 3  → worker 3 picks it up
// |--> sends job 4,5 → waiting in buffer
// |
// | workers start processing concurrently...
// | after 1s:
// |    worker 1 finishes job 1 → sends result → picks job 4
// |    worker 2 finishes job 2 → sends result → picks job 5
// |    worker 3 finishes job 3 → sends result → no job → exits
// |
// | after next 1s:
// |    worker 1 finishes job 4 → sends result → exits
// |    worker 2 finishes job 5 → sends result → exits
// |
// | wg.Wait() sees all Done → closes results channel
// | main() range results → prints all
// |
// 🏁 Program ends successfully


//////////////////////////////////////////////////////////////////////
// 💡 Key Mental Model
//////////////////////////////////////////////////////////////////////
//
// Component        | Role
// -----------------|-----------------------------------------------
// jobs channel     | A queue → distributes work to available goroutines
// workers          | Goroutines that take jobs, process, and send results
// results channel  | Collects outputs asynchronously
// WaitGroup        | Keeps main() waiting until all workers finish
// close()          | Used to signal “no more input/output”


// 🪄 Real-World Analogy
// ----------------------------------------------------
//
// 👨‍💼 Manager (main):
//   → Puts tasks into the task basket (jobs channel).
//
// 👷 Employees (workers):
//   → Pick up one task at a time, do the work, and drop the completed task
//     into the “done” basket (results channel).
//
// 🧾 Assistant (goroutine with WaitGroup):
//   → Waits quietly until all employees finish their work.
//     Once done, closes the office (results channel).
//
// 🏢 Manager (main) reads from the done basket and prints the results.
//
//////////////////////////////////////////////////////////////////////
