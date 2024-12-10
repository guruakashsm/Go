package main

import "fmt"

/*
=============================
DEFER IN GO
=============================

--- 1. Definition of `defer` ---
	`defer` is a Go keyword used to delay the execution of a function or statement until the surrounding function returns.
	Deferred calls are executed in **last in, first out** (LIFO) order.

--- 2. Purpose of `defer` ---
	`defer` is often used for tasks like closing resources, unlocking mutexes, or releasing file handles to ensure they are executed even if an error occurs.

--- 3. Execution Order ---
	Deferred calls are stacked and executed in reverse order of their appearance in the function.

--- 4. Arguments Evaluation ---
	Arguments to deferred function calls are evaluated immediately, but the function is executed later.

--- 5. Use Cases of `defer` ---
	1. Resource cleanup (e.g., closing files or database connections).
	2. Unlocking a locked resource (e.g., mutex).
	3. Simplifying code where cleanup should always be performed at the end of a function.

=============================
EXAMPLES OF `DEFER` IN GO
=============================

--- 6. Basic Usage of `defer` ---
	Using `defer` to execute a cleanup operation at the end of a function.

--- 7. Multiple `defer` Calls ---
	`defer` calls are executed in LIFO order (last in, first out).

--- 8. Arguments Evaluation ---
	Arguments for deferred calls are evaluated when the `defer` statement is executed, not when the deferred function is invoked.

*/

func main() {
	// Example 1: Basic Usage of `defer`
	fmt.Println("Start of main")
	defer fmt.Println("This will execute at the end of main")
	fmt.Println("End of main")

	// Example 2: Multiple `defer` Calls
	fmt.Println("\nMultiple defer calls example:")
	defer fmt.Println("First deferred call")
	defer fmt.Println("Second deferred call")
	defer fmt.Println("Third deferred call")

	// Example 3: Arguments Evaluation
	fmt.Println("\nArguments evaluation example:")
	value := 10
	defer fmt.Println("Deferred value:", value) // The value of `value` is captured now (10)
	value = 20                                  // Changing value will not affect the deferred call
	fmt.Println("Current value:", value)

	// Example 4: Using `defer` for Resource Cleanup
	fmt.Println("\nResource cleanup example:")
	resourceCleanup()
}

func resourceCleanup() {
	fmt.Println("Opening resource...")
	defer fmt.Println("Closing resource...") // Ensures resource is closed
	fmt.Println("Using resource...")
	// If an error occurs here, "Closing resource..." will still execute
}
