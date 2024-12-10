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
	`select` is a control flow keyword in Go that allows you to wait on multiple channels for a message or event. It is similar to a `switch` statement but for channels, enabling a program to handle multiple goroutines concurrently.

--- 2. Purpose of `select` ---
	`select` is used to:
	- Wait for multiple channel operations (sending or receiving) to complete.
	- Ensure non-blocking communication with multiple channels.
	- Handle timeouts or specific cases where multiple goroutines might be sending or receiving data concurrently.

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
	- `default`: Optional. Executes if no channel operation is ready.

--- 4. Blocking Behavior ---
	- By default, `select` blocks until one of the channels is ready for operation.
	- If multiple channels are ready, one is chosen at random to proceed.
	- If no channels are ready, and a `default` case is present, it will execute immediately without blocking.

--- 5. Use Cases ---
	- Handling multiple goroutines and their channel communication in a concurrent program.
	- Waiting for messages from multiple channels.
	- Implementing timeouts or default behavior when no channels are ready.
	- Coordinating between multiple goroutines to ensure proper synchronization.

=============================
EXAMPLES OF `SELECT` IN GO
=============================

--- 6. Basic `select` Usage ---
	Waiting on multiple channels and executing corresponding cases.

--- 7. `select` with Timeouts ---
	Using `select` to implement a timeout.

--- 8. `select` with Default Case ---
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
