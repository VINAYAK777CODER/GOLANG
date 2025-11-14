package main

import (
	"fmt"
	"sync"
)

// ❌ VERSION WITHOUT MUTEX → EXPLAINS PROBLEMS
// ------------------------------------------------
// Problem: Two goroutines modify the SAME variable (count)
// at the SAME time. This creates a "RACE CONDITION".
//
// What can go wrong?
// 1️⃣ Both goroutines read count at the same time (example: both read 0)
// 2️⃣ Both increment to 1
// 3️⃣ Both write 1 (one overwrites the other)
//     → final value should be 2, but becomes 1
//
// Output becomes unpredictable:
// Sometimes:
//   1 is the value
//   1 is the value
//
// Sometimes:
//   1 is the value
//   2 is the value
//
// Order keeps changing. Very dangerous.
// ------------------------------------------------

func increamenetNoMutex(count *int, wg *sync.WaitGroup) {
	defer wg.Done()

	// ❌ NOT SAFE: Read–Modify–Write without protection
	*count = *count + 1

	fmt.Println(*count, "is the value (NO MUTEX → unsafe)")
}

// ✅ VERSION WITH MUTEX (SAFE)
// ------------------------------------------------
// Mutex ensures ONLY ONE goroutine can access the shared
// variable at a time.
//
// Steps:
// 1️⃣ mu.Lock()  → take the lock (enter critical section)
// 2️⃣ modify count safely (no one else touching it now)
// 3️⃣ mu.Unlock() → release lock
//
// Now increments happen one by one, safely.
// ------------------------------------------------

func incrementWithMutex(count *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock() // 🚦 Only ONE goroutine can pass here at a time
	*count = *count + 1
	fmt.Println(*count, "is the value (WITH MUTEX → safe)")
	mu.Unlock() // 🔓 Release so next goroutine can continue
}

func main() {

	// -------------------------------
	// ❌ RUNNING UNSAFE VERSION
	// -------------------------------
	fmt.Println("\n--- Running WITHOUT mutex (unsafe) ---")

	{
		var wg sync.WaitGroup
		count := 0

		wg.Add(2)
		go increamenetNoMutex(&count, &wg)
		go increamenetNoMutex(&count, &wg)

		wg.Wait()

		fmt.Println("Final count without mutex (UNSAFE):", count)
	}

	// -------------------------------
	// ✅ RUNNING SAFE VERSION
	// -------------------------------
	fmt.Println("\n--- Running WITH mutex (safe) ---")

	{
		var wg sync.WaitGroup
		var mu sync.Mutex
		count := 0

		wg.Add(2)
		go incrementWithMutex(&count, &wg, &mu)
		go incrementWithMutex(&count, &wg, &mu)

		wg.Wait()

		fmt.Println("Final count with mutex (SAFE):", count)
	}

	fmt.Println("\nHello, World!")
}
