package main

import "fmt"

/*
=============================
FOR LOOP IN GO
=============================

--- 1. Definition of `for` ---
	`for` is the only looping construct in Go. It is used to repeatedly execute a block of code
	as long as a condition is satisfied or for a specific range of values.

--- 2. Purpose of `for` ---
	`for` allows you to iterate over:
	- Ranges (arrays, slices, maps, strings, etc.)
	- Numeric sequences with conditions or increments.
	- Infinite loops when no condition is specified.

--- 3. Syntax ---
	Basic structure of `for` loop:
	```go
	for initialization; condition; increment {
		// Code to execute repeatedly
	}
	```
	Range-based iteration:
	```go
	for index, value := range collection {
		// Code to process each element
	}
	```

--- 4. Variants of `for` ---
	- **Classic loop**: `for i := 0; i < 10; i++ { ... }`
	- **Range loop**: `for index, value := range collection { ... }`
	- **Infinite loop**: `for { ... }`
	- **Condition-only loop**: `for condition { ... }`

--- 5. Use Cases ---
	- Iterating over data structures like slices, maps, or strings.
	- Repeating tasks a fixed number of times.
	- Building custom loops with flexible conditions.

=============================
EXAMPLES OF `FOR` IN GO
=============================

--- 6. Classic `for` Loop ---
	Standard loop with initialization, condition, and increment.

--- 7. Iterating with `range` ---
	Using `range` to traverse slices, arrays, maps, and strings.

--- 8. Infinite Loop ---
	Creating an infinite loop for continuous tasks, terminated by `break`.

--- 9. Conditional Loop ---
	Looping with only a condition for more dynamic logic.

*/

func main() {
	// Example 1: Classic `for` loop
	fmt.Println("Classic `for` loop example:")
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// Example 2: Iterating with `range`
	fmt.Println("\n`for` with `range` example:")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Example 3: Infinite loop
	fmt.Println("\nInfinite `for` loop example:")
	counter := 0
	for {
		if counter == 3 {
			fmt.Println("Breaking the loop.")
			break
		}
		fmt.Println("Counter:", counter)
		counter++
	}

	// Example 4: Conditional loop
	fmt.Println("\nConditional loop example:")
	value := 10
	for value > 0 {
		fmt.Println("Value:", value)
		value -= 2
	}
}
