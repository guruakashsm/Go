package main

import "fmt"

/*
=============================
GO KEYWORD IN GO
=============================

--- 1. Definition of `go` ---
	`go` is a keyword in Go used to launch a new goroutine. A goroutine is a lightweight thread of execution,
	and the `go` keyword enables concurrent programming by running functions concurrently.

--- 2. Purpose of `go` ---
	`go` is used to execute a function asynchronously, meaning it allows functions to run concurrently
	with other functions, making it easier to perform tasks like I/O operations, network requests, or computations
	without blocking the main execution flow.

--- 3. Syntax ---
	The `go` keyword is used followed by the function call that you want to run concurrently:
	```go
	go functionName()
	```

--- 4. Goroutines ---
	A goroutine is a function that runs concurrently with other functions in Go. It runs in the background
	and does not block the main program, allowing for efficient use of CPU resources.

--- 5. Channel Synchronization ---
	When using goroutines, it is important to synchronize them with channels to ensure that the main program
	does not exit before all goroutines have finished executing.

--- 6. Use Cases ---
	- Performing background tasks asynchronously.
	- Handling multiple concurrent network requests.
	- Running long-running tasks in parallel to improve performance.

=============================
EXAMPLES OF `GO` IN GO
=============================

--- 7. Basic `go` usage ---
	Using the `go` keyword to launch a goroutine for concurrent execution.

--- 8. Synchronizing goroutines with channels ---
	Using channels to wait for goroutines to finish execution.

*/

func main() {
	// Example 1: Basic `go` usage
	fmt.Println("Basic `go` example:")
	go greet() // This runs concurrently

	// Wait for the goroutine to finish before the main function exits
	// Sleep is just to simulate the wait (for demo purposes)
	// In production, you would typically use channels or other synchronization methods
	fmt.Println("Main function continues executing...")

	// Example 2: Synchronizing goroutines with channels
	fmt.Println("\nSynchronizing goroutines example:")
	done := make(chan bool) // Channel to synchronize goroutines

	go asyncTask(done) // Start the goroutine
	<-done             // Wait for the goroutine to signal that it has finished

	fmt.Println("Goroutine has finished execution.")
}

// Example 1: Basic function declaration with go keyword
func greet() {
	fmt.Println("Hello from a goroutine!")
}

// Example 2: Function to simulate an asynchronous task
func asyncTask(done chan bool) {
	// Simulating some task with a simple print statement
	fmt.Println("Performing an asynchronous task...")
	done <- true // Signal that the task is done
}
