package main

import (
	"fmt"
	"time"
)

/*
=============================
SELECT IN GO
=============================

--- 1. Definition of `select` ---
	`select` is a control flow keyword in Go that is used for waiting on multiple channels. It allows you to simultaneously wait for data from multiple channels and act accordingly when data is available. Unlike a `switch`, which evaluates values, `select` works with channels and allows the program to handle concurrent operations.

--- 2. Purpose of `select` ---
	`select` is used to:
	- Wait for multiple channels to become ready for communication.
	- Handle multiple goroutines concurrently by selecting the first channel that is ready.
	- Implement non-blocking behavior by using the `default` case when no channels are ready.
	- Implement timeouts or delay handling with `time.After`.

--- 3. Syntax of `select` ---
	```go
	select {
	case <-channel1:
		// Handle message from channel1
	case <-channel2:
		// Handle message from channel2
	case <-channel3:
		// Handle message from channel3
	default:
		// Handle default case if no channel operation is ready
	}
	```

	- `case`: Defines a channel operation to wait for.
	- `default`: Executes when no channel operation is ready.

--- 4. Blocking Behavior ---
	- By default, `select` blocks until at least one of the channels is ready.
	- If multiple channels are ready, one is chosen randomly to execute.
	- If no channels are ready, and a `default` case is present, it executes immediately without blocking.

--- 5. Difference Between `select` and `switch` ---
	- **`select`**: Designed for working with channels. It blocks until one of the channels is ready, allowing concurrent handling of multiple communication channels. `select` is particularly useful in concurrent programming when you need to wait for messages from multiple goroutines.
	- **`switch`**: A traditional conditional statement used for comparing values. It evaluates expressions and chooses the case that matches the value, but it does not handle channels or concurrency directly. `switch` can be used for simple branching based on value comparisons.

	Example:
	```go
	// `switch` for value comparison
	switch value {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	}

	// `select` for channel operations
	select {
	case msg := <-ch1:
		fmt.Println("Received from ch1:", msg)
	case msg := <-ch2:
		fmt.Println("Received from ch2:", msg)
	}
	```

--- 6. Use Cases ---
	- Handling communication between multiple goroutines via channels.
	- Waiting for messages from multiple sources and acting on the first one ready.
	- Implementing timeouts using `select` with `time.After()`.
	- Implementing default behavior when no channels are ready for communication.

=============================
EXAMPLES OF `SELECT` IN GO
=============================

--- 7. Basic `select` Usage ---
	Waiting on multiple channels and executing corresponding cases.

--- 8. `select` with Timeouts ---
	Using `select` to implement a timeout.

--- 9. `select` with Default Case ---
	Handling a case when no channels are ready.

*/

func main() {
	// Example 1: Basic `select` Usage
	fmt.Println("Basic `select` example:")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	}

	// Example 2: `select` with Timeouts
	fmt.Println("\n`select` with timeout example:")
	timeout := time.After(3 * time.Second)
	select {
	case msg := <-ch1:
		fmt.Println("Received:", msg)
	case msg := <-ch2:
		fmt.Println("Received:", msg)
	case <-timeout:
		fmt.Println("Timeout occurred!")
	}

	// Example 3: `select` with Default Case
	fmt.Println("\n`select` with default case example:")
	select {
	case msg := <-ch1:
		fmt.Println("Received:", msg)
	case msg := <-ch2:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message received, proceeding with default case.")
	}
}
