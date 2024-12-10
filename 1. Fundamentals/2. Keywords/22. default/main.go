package main

import (
	"fmt"
	"log"
	"time"
)

/*
=============================
DEFAULT IN GO
=============================

--- 1. Definition of `default` ---
	The `default` keyword is used in a `select` statement in Go to specify a case that will execute if no other channel is ready for communication.
	It allows the program to avoid blocking when no other cases are available to proceed.

--- 2. How `default` Works ---
	When used in a `select` statement, the `default` case will execute immediately if none of the channels are ready for communication.
	It helps to provide non-blocking behavior when you're working with multiple channels.

--- 3. Use Case of `default` ---
	`default` is commonly used when you want to handle a situation where no channels are ready or you want to implement a timeout mechanism without blocking forever.

=============================
EXAMPLES OF `DEFAULT` IN GO
=============================

--- 4. Basic Usage of `default` in a `select` ---
	Using `default` to avoid blocking when no channels are ready.

--- 5. `default` for Non-blocking `select` ---
	Using `default` to prevent blocking the `select` statement when no case is ready.

--- 6. `default` with Timeout ---
	Using `default` to implement a timeout mechanism in the `select` statement.

*/

func main() {
	// Example 1: Basic Usage of default in select
	ch1 := make(chan int)
	go func() {
		ch1 <- 42
	}()
	select {
	case msg := <-ch1:
		fmt.Println("Received from ch1:", msg)
	default:
		fmt.Println("No message received, executing default case.")
	}

	// Example 2: default for Non-blocking select
	ch2 := make(chan int)
	go func() {
		// No message sent to ch2
	}()
	select {
	case msg := <-ch2:
		fmt.Println("Received from ch2:", msg)
	default:
		fmt.Println("ch2 is empty, executing default case.")
	}

	// Example 3: default with Timeout
	ch3 := make(chan int)
	select {
	case msg := <-ch3:
		fmt.Println("Received from ch3:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout occurred, executing default case.")
	}

	// Example 4: Using default with switch
	switch "MONDAY" {
	case "SUNDAY":
		log.Println("It's Sunday!")
	default:
		log.Println("It's not Sunday!")
	}

	// Example 5: Using default with switch for type assertion
	var val interface{} = 42
	switch val.(type) {
	case int:
		log.Println("Value is an integer.")
	default:
		log.Println("Value is not an integer.")
	}
}
