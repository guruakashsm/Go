package main

import "fmt"

/*
=============================
RETURN IN GO
=============================

--- 1. Definition of `return` ---
	The `return` keyword is used to exit from a function and optionally pass values back to the caller. It allows a function to send a result (or multiple results) back after performing its logic.

--- 2. Purpose of `return` ---
	`return` is used to:
	- Exit the function at a specific point.
	- Send back values to the caller (if the function is designed to return any values).
	- Terminate the function's execution, stopping further code execution inside the function.

--- 3. Syntax of `return` ---
	```go
	return
	```
	To return values from a function, you can specify the values after the `return` keyword:
	```go
	return value1, value2
	```

	- `value1`, `value2`: The values to be returned from the function. These values must match the function's return type.

--- 4. Return with No Values ---
	Functions in Go can have no return values. In this case, `return` is used simply to exit the function without sending anything back.

--- 5. Return with Multiple Values ---
	Go allows functions to return multiple values. This is commonly used for returning both a result and an error in many standard library functions.

--- 6. Use Cases ---
	- Exiting a function early when a specific condition is met.
	- Returning calculated values from a function.
	- Returning multiple values, such as data and error, from a function.
	- Stopping function execution before the end of its logic.

=============================
EXAMPLES OF `RETURN` IN GO
=============================

--- 7. Basic `return` Usage ---
	Using `return` to send a value back to the caller from a function.

--- 8. Returning Multiple Values ---
	Returning multiple values from a function, often used for error handling.

--- 9. Using `return` to Exit Early ---
	Exiting a function early using `return` when a condition is met.

*/

func main() {
	// Example 1: Basic `return` Usage
	fmt.Println("Basic `return` example:")
	result := add(3, 5)
	fmt.Println("Result of addition:", result)

	// Example 2: Returning Multiple Values
	fmt.Println("\nReturning Multiple Values:")
	data, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result of division:", data)
	}

	// Example 3: Using `return` to Exit Early
	fmt.Println("\nExit Early with `return`:")
	earlyExit(10)
}

// Function with a single return value
func add(a, b int) int {
	return a + b
}

// Function with multiple return values (result and error)
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Function to demonstrate early return
func earlyExit(num int) {
	if num > 5 {
		fmt.Println("Exiting function early as num > 5")
		return // Exits the function early
	}
	fmt.Println("Continuing execution as num <= 5")
}
